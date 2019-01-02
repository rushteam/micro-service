package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gitee.com/rushteam/micro-service/common/sdk/wxsdk/wxpay"
	"gitee.com/rushteam/micro-service/common/utils"
	"gitee.com/rushteam/micro-service/service/pay-srv/config"
	"gitee.com/rushteam/micro-service/service/pay-srv/model"
	"gitee.com/rushteam/micro-service/service/pay-srv/queue"
	"github.com/micro/go-micro/errors"

	// "gitee.com/rushteam/micro-service/common/utils"
	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	log "github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/pborman/uuid"
	// "go.uber.org/zap"
)

const (
	//TradeTypeWxJsAPI ...
	TradeTypeWxJsAPI = "JSAPI"
	//TradeWxpay ..
	TradeWxpay = "wxpay"
	//TradeAlipay ..
	TradeAlipay = "alipay"
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
	if req.GetAccessToken() == "" {
		return errors.BadRequest("PayService.Create", "params err, client_id is undefined")
	}
	//todo check token
	//permission denied

	if req.GetChannel() == "" {
		return errors.BadRequest("PayService.Create", "params err, channel is undefined")
	}
	if req.GetTotalFee() <= 0 {
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
	clientID := req.GetClientId()
	clientConf, ok := config.App.Apps[clientID]
	if !ok {
		return errors.BadRequest("PayService.Create", "not found client_id: "+clientID)
	}
	//todo 检测商户秘钥是否正确

	payChannelID := req.GetChannel()
	var isAblePayChannel = false
	for _, v := range clientConf.PayChannels {
		if payChannelID == v {
			isAblePayChannel = true
			break
		}
	}
	if isAblePayChannel == false {
		return errors.BadRequest("PayService.Create", "not found channel for pay in this client: "+payChannelID)
	}
	payConf, ok := config.App.PayChannels[payChannelID]
	if !ok {
		return errors.BadRequest("PayService.Create", "not found channel for pay: "+payChannelID)
	}
	if payConf.MchID == "" {
		return errors.BadRequest("PayService.Create", fmt.Sprintf("channel(%s) info is incomplete", payChannelID))
	}
	if payConf.AppID == "" {
		return errors.BadRequest("PayService.Create", fmt.Sprintf("channel(%s) info is incomplete", payChannelID))
	}
	if payConf.NotifyURL == "" {
		return errors.BadRequest("PayService.Create", fmt.Sprintf("channel(%s) info is incomplete", payChannelID))
	}

	//生成 订单信息
	tradeModel := model.TradeModel{}
	tradeModel.ClientId = clientID
	tradeModel.PvdMchId = payConf.MchID
	tradeModel.PvdAppid = payConf.AppID
	tradeModel.Channel = payChannelID

	tradeModel.OutTradeNo = req.GetOutTradeNo()
	tradeModel.TotalFee = req.GetTotalFee()
	tradeModel.Subject = req.GetSubject()
	tradeModel.FromIp = req.GetFromIp()
	tradeModel.TradeType = req.GetTradeType()
	tradeModel.PvdNotifyUrl = payConf.NotifyURL
	tradeModel.PvdOutTradeNo = tradeModel.OutTradeNo //暂时透传 应用方的第三方单号
	// tradeModel.PvdTradeNo = ""                       //调用第三方支付成功后赋值
	tradeModel.Provider = payConf.Provider

	if tradeModel.Provider == TradeWxpay { //微信
		order := &wxpay.UnifiedOrderReq{}
		order.AppID = tradeModel.PvdAppid
		order.MchID = tradeModel.PvdMchId

		order.OutTradeNo = tradeModel.PvdOutTradeNo //商户订单号
		order.TotalFee = tradeModel.TotalFee        //订单总金额，单位为分
		order.FeeType = "RMB"                       //标价币种 目前写死
		order.Body = tradeModel.Subject             //商品描述 128
		order.NotifyURL = payConf.NotifyURL         //异步通知地址
		order.TradeType = tradeModel.TradeType      //TradeType  (JSAPI|NATIVE)

		if tradeModel.TradeType == TradeTypeWxJsAPI {
			//仅在 TradeType=JSAPI 时必须
			extra := make(map[string]string)
			rawExtra := []byte(req.GetExtra())
			json.Unmarshal(rawExtra, &extra)
			if _, ok := extra["openid"]; !ok {
				return errors.BadRequest("PayService.Create", "params err, openid is undefined")
			}
			order.OpenID = extra["openid"]
		}

		// order.OpenID = "o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ"
		order.SpbillCreateIP = tradeModel.FromIp
		order.NonceStr = utils.RandomStr(32) //随机字符串
		order.MakeSign(payConf.ApiKey)
		orderRsp, err := order.Call()
		if err != nil {
			return errors.BadRequest("PayService.Create", "pay channel call err, "+err.Error())
		}
		if err = orderRsp.Error(); err != nil {
			return errors.BadRequest("PayService.Create", "pay channel resp err, "+err.Error())
		}
		//赋值第三方交易号
		tradeModel.PvdTradeNo = orderRsp.PrepayID
		//rsp.ClientId = tradeModel.ClientId
		payField := &pay_srv.PayField{
			AppId:      tradeModel.PvdAppid,
			OutTradeNo: tradeModel.PvdOutTradeNo,
			TradeNo:    tradeModel.PvdTradeNo,
			TotalFee:   tradeModel.TotalFee,
			FieldStr:   "",
		}
		rsp.Channel = tradeModel.Channel
		rsp.Provider = tradeModel.Provider
		if tradeModel.TradeType == TradeTypeWxJsAPI {
			payConfig := &wxpay.PayConfigJs{
				AppID:     order.AppID,
				TimeStamp: time.Now().Unix(),
				NonceStr:  utils.RandomStr(32),
				Package:   fmt.Sprintf("prepay_id=%s", orderRsp.PrepayID),
			}
			payConfig.MakeSign(payConf.ApiKey)
			jsonBytes, err := json.Marshal(payConfig)
			if err != nil {
				return errors.BadRequest("PayService.Create", "pay channel jsapi err, "+err.Error())
			}
			payField.FieldStr = string(jsonBytes)
		}
		rsp.PayField = payField
	} else if tradeModel.Provider == TradeAlipay {
		return errors.BadRequest("PayService.Create", "alipay channel is undefined")
	} else {
		return errors.BadRequest("PayService.Create", "pay channel is undefined")
	}
	//保存订单到数据库
	return nil
}

//Notify ..
func (s *PayService) Notify(ctx context.Context, req *pay_srv.NotifyReq, rsp *pay_srv.NotifyRsp) error {
	if req.GetChannel() == "" {
		return errors.BadRequest("PayService.Notify", "params err, channel is undefined")
	}
	if req.GetRaw() == "" {
		return errors.BadRequest("PayService.Notify", "params err, raw is undefined")
	}
	payChannelID := req.GetChannel()
	payConf, ok := config.App.PayChannels[payChannelID]
	if !ok {
		return errors.BadRequest("PayService.Notify", "not found channel for pay: "+payChannelID)
	}
	raw := req.GetRaw()
	if payConf.Provider == TradeWxpay { //微信支付
		notify, err := wxpay.UnmarshalNotify(raw)
		if err != nil {
			return errors.BadRequest("PayService.Notify", "params err, raw is invalid")
		}
		if notify.IsSuccess() == false {
			rsp.Result = wxpay.NotifyReplyFail("服务商返回失败")
			return nil
		}
		if wxpay.CheckSign(payConf.ApiKey, notify) == false {
			rsp.Result = wxpay.NotifyReplyFail("签名校验失败")
			return nil
		}
		// utils.FormatDate(time.Now()),
		//修改状态
		// tm := model.TradeModel{
		// 	PayState: 1,
		// 	PayAt:    time.Now(),
		// }
		// err = tm.Save()
		// if err != nil {
		// 	rsp.Result = wxpay.NotifyReplyFail("交易数据存储时发生错误")
		// 	return errors.BadRequest("PayService.Notify", "trade data save fail")
		// }
		//进行真实回调任务
		// queue.PayNotify.Publish(&pay_srv.NotifyApp{})
		//进行真实回调任务
		ev := &pay_srv.NotifyEvent{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Name:      "this first msg in system",
			Url:       "https://1thx.com",
			Body:      "{}",
		}
		err = queue.Publish(ctx, "pay_notify", ev)
		if err != nil {
			log.Fatal(err)
		}
		//返回支付成功信息
		rsp.Result = wxpay.NotifyReplySuccess()
		rsp.OutTradeNo = notify.OutTradeNo
		//支付成功后
		// } else if payConf.Provider == TradeAlipay { //阿里支付
	} else {
		return errors.BadRequest("PayService.Create", "pay channel is undefined")
	}
	return nil
}

//Query ..
func (s *PayService) Query(ctx context.Context, req *pay_srv.QueryReq, rsp *pay_srv.PayRsp) error {
	// log.Log("[access] PayService.Query")
	return nil
}
