package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

//PayService ...
type PayService struct {
	Service micro.Service
	// logger *zap.Logger
}

//validateCreateReq
func validateCreateReq(req *pay_srv.CreateReq) error {
	if req.GetOutTradeNo() == "" {
		return errors.BadRequest("PayService.Create", "参数不全,缺少out_trade_no")
	}
	if req.GetClientId() == "" {
		return errors.BadRequest("PayService.Create", "参数不全,缺少client_id")
	}
	if req.GetAccessToken()== "" /** || token检测*/{
		return errors.BadRequest("PayService.Create", "权限不足")
	}
	if req.GetChannel() == "" {
		return errors.BadRequest("PayService.Create", "参数不全,缺少channel")
	}
	if req.GetTotalFee() <=0 {
		return errors.BadRequest("PayService.Create", "支付金额不能小于零")
	}
	if req.GetSubject() == "" {
		return errors.BadRequest("PayService.Create", "参数不全,缺少subject")
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
	//用来覆盖默认设置
	if req.GetTradeType() != "" {
		//支付方式
		app.TradeType = req.GetTradeType()
	}
	//生成 订单信息
	tradeModel := model.TradeModel{}
	tradeNo := md5.New()
	tradeNo.Write([]byte(req.GetOutTradeNo()))
	tradeModel.TradeNo = hex.EncodeToString(tradeNo.Sum(nil))
	tradeModel.OutTradeNo = req.GetOutTradeNo()
	tradeModel.TotalFee = req.GetTotalFee()
	tradeModel.Subject = req.GetSubject()
	tradeModel.Channel = req.GetChannel()
	tradeModel.ProviderName = app.Channel
	tradeModel.ProviderMchId = app.MchID
	tradeModel.ProviderAppid = app.AppID
	tradeModel.FromIp = req.GetFromIp()
	tradeModel.ClientId = req.ClientId
	//微信
	if app.Channel == "wxpay" {
		order := &mch.UnifiedOrderReq{}
		order.AppID = tradeModel.ProviderAppid
		order.MchID = tradeModel.ProviderMchId

		order.OutTradeNo = tradeModel.TradeNo //商户订单号
		order.TotalFee = tradeModel.TotalFee     //订单总金额，单位为分
		order.FeeType = "RMB"                  //标价币种 目前写死
		order.Body = tradeModel.Subject          //商品描述 128
		order.NotifyURL = "https://test.com"   //异步通知地址
		order.TradeType = app.TradeType        //TradeType
		order.OpenID = req.GetOpenId() //仅在 TradeType=JSAPI 时必须
		// order.OpenID = "o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ"
		order.SpbillCreateIP = tradeModel.FromIp
		order.NonceStr = utils.RandomStr(32) //随机字符串
		order.MakeSign(app.ApiKey)
		orderRsp, err := order.Call()
		if err != nil {
			return errors.BadRequest("PayService.Create", "pay channel call err: " + err.Error())
		}
		if err = orderRsp.Error(); err != nil {
			return errors.BadRequest("PayService.Create", "pay channel resp err: " + err.Error())
		}
		//orderRsp.PrepayID
		rsp.ClientId = tradeModel.ClientId
		rsp.OutTradeNo = tradeModel.OutTradeNo
		rsp.Channel = tradeModel.Channel
		rsp.TradeNo = tradeModel.TradeNo
		if app.TradeType == "JSAPI" {
			payConfig := &mch.PayConfigJs{
				AppID:     order.AppID,
				TimeStamp: time.Now().Unix(),
				NonceStr:  utils.RandomStr(32),
				Package:   fmt.Sprintf("prepay_id=%s", orderRsp.PrepayID),
				SignType:  signType,
			}
			payConfig.MakeSign(app.ApiKey)
			payField, err := json.Marshal(payConfig)
			if err != nil {
				return errors.BadRequest("PayService.Create", "pay channel jsapi err: " + err.Error())
			}
			rsp.PayField = string(payField)
		}

	}
	//创建订单

	//赋值订单
	rsp.OutTradeNo = req.GetOutTradeNo()
	rsp.Channel = req.GetChannel()
	rsp.TotalFee = req.GetTotalFee()
	return nil
}

//Notify ..
func (s *PayService) Notify(ctx context.Context, req *pay_srv.NotifyReq, rsp *pay_srv.PayRsp) error {
	log.Log("[access] PayService.Notify")
	return nil
}

//Query ..
func (s *PayService) Query(ctx context.Context, req *pay_srv.QueryReq, rsp *pay_srv.PayRsp) error {
	log.Log("[access] PayService.Query")
	return nil
}
