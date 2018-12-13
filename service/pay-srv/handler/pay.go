package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/rushteam/micro-service/common/sdk/wxsdk"
	"gitee.com/rushteam/micro-service/common/sdk/wxsdk/mch"
	"gitee.com/rushteam/micro-service/common/utils"
	"gitee.com/rushteam/micro-service/service/pay-srv/config"
	"gitee.com/rushteam/micro-service/service/pay-srv/model"
	"github.com/micro/go-micro/errors"
	"time"

	// "gitee.com/rushteam/micro-service/common/utils"
	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	// "go.uber.org/zap"
)
const(
	TradeTypeWxJsApi = "WX_JSAPI"
)
//PayService ...
type PayService struct {
	Service micro.Service
	// logger *zap.Logger
}

//validateCreateReq
func validateCreateReq(req *pay_srv.CreateReq) error {
	if req.GetOutTradeNo() == "" {
		return errors.BadRequest("PayService.Create", "params err, out_trade_no is undefined")
	}
	if req.GetClientId() == "" {
		return errors.BadRequest("PayService.Create", "params err, client_id is undefined")
	}
	if req.GetAccessToken()== ""{
		return errors.BadRequest("PayService.Create", "params err, client_id is undefined")
	}
	//todo check token
	//permission denied

	if req.GetChannel() == "" {
		return errors.BadRequest("PayService.Create", "params err, channel is undefined")
	}
	if req.GetTotalFee() <=0 {
		return errors.BadRequest("PayService.Create", "params err, total_fee must not be less then 0")
	}
	if req.GetSubject() == "" {
		return errors.BadRequest("PayService.Create", "params err, subject is undefined")
	}
	return nil
}
//Create ..
func (s *PayService) Create(ctx context.Context, req *pay_srv.CreateReq, rsp *pay_srv.PayRsp) error {
	err := validateCreateReq(req)
	if err != nil {
		return err
	}
	//创建订单
	var signType = "MD5"
	clientID := req.GetClientId()
	clientInfo, ok := config.App.Apps[clientID]
	if !ok {
		return errors.BadRequest("PayService.Create", "not found client_id: " + clientID)
	}
	payChannel := req.GetChannel()
	app, ok := clientInfo.Channels[payChannel]
	if !ok {
		return errors.BadRequest("PayService.Create", "not found channel for pay: " + req.GetChannel())
	}

	//生成 订单信息
	tradeModel := model.TradeModel{}
	tradeModel.ClientId = clientID
	tradeModel.PvdMchId = app.MchID
	tradeModel.PvdAppid = app.AppID
	tradeModel.Channel = req.GetChannel()
	tradeModel.OutTradeNo = req.GetOutTradeNo()
	tradeModel.TotalFee = req.GetTotalFee()
	tradeModel.Subject = req.GetSubject()
	tradeModel.FromIp = req.GetFromIp()
	tradeModel.TradeType = req.GetTradeType()
	tradeModel.PvdOutTradeNo = tradeModel.OutTradeNo //暂时透传 应用方的第三方单号
	tradeModel.PvdTradeNo = "" //调用第三方支付成功后赋值

	tradeModel.ProviderName = app.Channel

	if app.Channel == "wxpay" { //微信
		order := &mch.UnifiedOrderReq{}
		order.AppID = tradeModel.PvdAppid
		order.MchID = tradeModel.PvdMchId

		order.OutTradeNo = tradeModel.OutTradeNo //商户订单号
		order.TotalFee = tradeModel.TotalFee     //订单总金额，单位为分
		order.FeeType = "RMB"                  //标价币种 目前写死
		order.Body = tradeModel.Subject          //商品描述 128
		order.NotifyURL = "https://test.com"   //异步通知地址
		order.TradeType = tradeModel.TradeType       //TradeType  (JSAPI|NATIVE)

		if tradeModel.TradeType == "JSAPI" {
			//仅在 TradeType=JSAPI 时必须
			extra := make(map[string]string)
			rawExtra := []byte(req.GetExtra())
			json.Unmarshal(rawExtra,&extra)
			if _,ok := extra["openid"];!ok {
				return errors.BadRequest("PayService.Create", "params err, openid is undefined")
			}
			order.OpenID = extra["openid"]
		}

		// order.OpenID = "o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ"
		order.SpbillCreateIP = tradeModel.FromIp
		order.NonceStr = utils.RandomStr(32) //随机字符串
		order.MakeSign(app.ApiKey)
		orderRsp, err := order.Call()
		if err != nil {
			return errors.BadRequest("PayService.Create", "pay channel call err, " + err.Error())
		}
		if err = orderRsp.Error(); err != nil {
			return errors.BadRequest("PayService.Create", "pay channel resp err, " + err.Error())
		}
		//赋值第三方交易号
		tradeModel.PvdTradeNo = orderRsp.PrepayID
		//rsp.ClientId = tradeModel.ClientId
		payField := &pay_srv.PayField{
			AppId:	tradeModel.PvdAppid,
			OutTradeNo: tradeModel.OutTradeNo,
			TradeNo: tradeModel.TradeNo,
			TotalFee: tradeModel.TotalFee,
			FieldStr: "",
		}
		rsp.Channel = tradeModel.Channel
		rsp.ProviderName = tradeModel.ProviderName
		if tradeModel.TradeType == "JSAPI" {
			payConfig := &mch.PayConfigJs{
				AppID:     order.AppID,
				TimeStamp: time.Now().Unix(),
				NonceStr:  utils.RandomStr(32),
				Package:   fmt.Sprintf("prepay_id=%s", orderRsp.PrepayID),
				SignType:  signType,
			}
			payConfig.MakeSign(app.ApiKey)
			jsonBytes, err := json.Marshal(payConfig)
			if err != nil {
				return errors.BadRequest("PayService.Create", "pay channel jsapi err, " + err.Error())
			}
			payField.FieldStr = string(jsonBytes)
		}
		rsp.PayField = payField
	}else if app.Channel == "alipay" {
		return errors.BadRequest("PayService.Create", "alipay channel is undefined")
	} else {
		return errors.BadRequest("PayService.Create", "pay channel is undefined")
	}
	//保存订单到数据库
	return nil
}

//Notify ..
func (s *PayService) Notify(ctx context.Context, req *pay_srv.NotifyReq, rsp *pay_srv.PayRsp) error {
	log.Log("[access] PayService.Notify")
	if req.GetPvdName() == "" {
		return errors.BadRequest("PayService.Notify", "params err, pvd_name is undefined")
	}
	if req.GetPvdName() == "wxpay" { //微信支付
		if req.GetRaw() == "" {
			return errors.BadRequest("PayService.Notify", "params err, raw is undefined")
		}
		raw := req.GetRaw()
		notify,err := mch.UnmarshalNotify(raw)
		if err != nil {
			return errors.BadRequest("PayService.Notify", "params err, raw is invalid")
		}
		fmt.Println(notify.OutTradeNo)
	} else if req.GetPvdName() == "alipay" { //阿里支付

	} else {
		return errors.BadRequest("PayService.Create", "pay channel is undefined")
	}
	return nil
}

//Query ..
func (s *PayService) Query(ctx context.Context, req *pay_srv.QueryReq, rsp *pay_srv.PayRsp) error {
	log.Log("[access] PayService.Query")
	return nil
}
