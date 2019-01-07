package handler

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"

	"gitee.com/rushteam/micro-service/common/utils/snowflake"

	"github.com/pborman/uuid"

	"github.com/mlboy/godb/orm"

	"gitee.com/rushteam/micro-service/common/sdk/wxsdk/wxpay"
	"gitee.com/rushteam/micro-service/common/utils"
	"gitee.com/rushteam/micro-service/service/pay-srv/config"
	"gitee.com/rushteam/micro-service/service/pay-srv/model"
	"gitee.com/rushteam/micro-service/service/pay-srv/queue"
	"github.com/micro/go-micro/errors"

	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	log "github.com/micro/go-log"
	micro "github.com/micro/go-micro"
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
	if req.GetOutPayNo() == "" {
		return errors.BadRequest("PayService.Create", "params err, out_pay_no is undefined")
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
	if req.GetNotifyUrl() == "" {
		return errors.BadRequest("PayService.Create", "params err, notify_url is undefined")
	}
	return nil
}

//Create ..
func (s *PayService) Create(ctx context.Context, req *pay_srv.CreateReq, rsp *pay_srv.PayRsp) error {
	// fmt.Println(s.Service.Server().Options().Registry.GetService("go.micro.srv.pay_srv"))
	// md, ok := metadata.FromContext(ctx)
	// if !ok {
	// 	md = metadata.Metadata{}
	// }
	// fmt.Println(md)
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
	//TODO: 检测商户秘钥是否正确

	// err = checkSign(req, "secret")
	// if err != nil {
	// 	fmt.Println(err)
	// }

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

	//生成单号
	sn, _ := snowflake.NewSnowFlake(1)
	payNo := strconv.FormatUint(sn.Next(), 10)
	if payNo == "" {
		return errors.BadRequest("PayService.Create", "not created payNo ")
	}
	//生成 订单信息
	tradeModel := &model.TradeModel{}
	tradeModel.PayNo = payNo //第三方单号 使用网关支付号
	tradeModel.ClientID = clientID
	tradeModel.PvdMchID = payConf.MchID
	tradeModel.PvdAppID = payConf.AppID
	tradeModel.Channel = payChannelID

	tradeModel.OutPayNo = req.GetOutPayNo()
	tradeModel.TotalFee = req.GetTotalFee()
	tradeModel.Subject = req.GetSubject()
	tradeModel.NotifyURL = req.GetNotifyUrl()
	tradeModel.FromIP = req.GetFromIp()
	tradeModel.TradeType = req.GetTradeType()
	tradeModel.PvdNotifyURL = payConf.NotifyURL
	// tradeModel.PvdTradeNo = ""                       //调用第三方支付成功后赋值
	tradeModel.Provider = payConf.Provider
	tradeModel.FeeType = "RMB"

	//保存订单到数据库
	_, err = orm.Model(tradeModel).Insert()
	if err != nil {
		if me, ok := err.(*mysql.MySQLError); ok {
			if me.Number == 1062 {
				return errors.BadRequest("PayService.Create", "支付单号%s已存在", tradeModel.OutPayNo)
			}
		}
		return errors.BadRequest("PayService.Create", "insert trade record error "+err.Error())
	}
	if tradeModel.Provider == TradeWxpay { //微信
		order := &wxpay.UnifiedOrder{}
		order.AppID = tradeModel.PvdAppID
		order.MchID = tradeModel.PvdMchID

		order.OutTradeNo = payNo               //商户订单号
		order.TotalFee = tradeModel.TotalFee   //订单总金额，单位为分
		order.FeeType = tradeModel.FeeType     //标价币种 目前写死
		order.Body = tradeModel.Subject        //商品描述 128
		order.NotifyURL = payConf.NotifyURL    //异步通知地址
		order.TradeType = tradeModel.TradeType //TradeType  (JSAPI|NATIVE)
		//写死md5支付
		order.SignType = "MD5"
		var signHandler = md5.New()

		if tradeModel.TradeType == TradeTypeWxJsAPI {
			//仅在 TradeType=JSAPI 时必须
			extra := make(map[string]string)
			rawExtra := []byte(req.GetExtra())
			json.Unmarshal(rawExtra, &extra)
			if _, ok := extra["openid"]; !ok {
				return errors.BadRequest("PayService.Create", "params error: openid is undefined")
			}
			order.OpenID = extra["openid"]
		}

		// order.OpenID = "o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ"
		order.SpbillCreateIP = tradeModel.FromIP
		order.NonceStr = utils.RandomStr(32) //随机字符串
		order.Sign = wxpay.Sign(order, payConf.ApiKey, signHandler)
		orderRsp := &wxpay.UnifiedOrderRsp{}
		err := wxpay.Request(order, orderRsp)
		if err != nil {
			return errors.BadRequest("PayService.Create", "pay channel error: %s", err.Error())
		}
		//赋值第三方交易号
		tradeModel.PvdTradeNo = orderRsp.PrepayID
		//rsp.ClientId = tradeModel.ClientId
		payField := &pay_srv.PayField{
			AppId:      tradeModel.PvdAppID,
			OutTradeNo: payNo,
			TradeNo:    tradeModel.PvdTradeNo,
			TotalFee:   tradeModel.TotalFee,
			FieldStr:   "",
		}
		rsp.Channel = tradeModel.Channel
		rsp.Provider = tradeModel.Provider
		if tradeModel.TradeType == TradeTypeWxJsAPI {
			payment := &wxpay.Payment{
				AppID:     order.AppID,
				TimeStamp: time.Now().Unix(),
				NonceStr:  utils.RandomStr(32),
				Package:   fmt.Sprintf("prepay_id=%s", orderRsp.PrepayID),
			}
			payment.PaySign = wxpay.Sign(payment, payConf.ApiKey, signHandler)
			paymentBytes, err := json.Marshal(payment)
			if err != nil {
				return errors.BadRequest("PayService.Create", "pay channel jsapi err, "+err.Error())
			}
			payField.FieldStr = string(paymentBytes)
		}
		rsp.PayField = payField
	} else if tradeModel.Provider == TradeAlipay {
		return errors.BadRequest("PayService.Create", "alipay channel is undefined")
	} else {
		return errors.BadRequest("PayService.Create", "pay channel is undefined")
	}
	//保存订单到数据库
	_, err = orm.Model(tradeModel).Update()
	if err != nil {
		return errors.BadRequest("PayService.Create", "save trade record error")
	}
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
		notify := &wxpay.Notify{}
		err := wxpay.Response([]byte(raw), notify)
		if err != nil {
			return errors.BadRequest("PayService.Notify", "params err, raw is invalid")
		}
		var signHandler hash.Hash
		if notify.SignType == "MD5" {
			signHandler = md5.New()
		} else {
			signHandler = md5.New()
		}
		if wxpay.Sign(notify, payConf.ApiKey, signHandler) != notify.Sign {
			rsp.Result = wxpay.NotifyReplyFail("sign error")
			return nil
		}
		if notify.OutTradeNo == "" {
			rsp.Result = wxpay.NotifyReplyFail("缺少订单号数据")
			return nil
		}
		//支付成功后
		//查找订单
		tm := &model.TradeModel{}
		err = orm.Model(tm).Where("pay_no", notify.OutTradeNo).Find()
		if err != nil {
			log.Logf("PayService.Notify not_found_trade_record %+v", notify)
			return errors.BadRequest("PayService.Notify", fmt.Sprintf("not found trade record, pay_no=%s", notify.OutTradeNo))
		}
		// utils.FormatDate(time.Now()),
		//修改状态
		tm.PayState = 1
		tm.PayAt = time.Now()
		_, err = orm.Model(tm).Where("pay_no", notify.OutTradeNo).Update()
		if err != nil {
			rsp.Result = wxpay.NotifyReplyFail("存储交易数据时发生错误")
			return errors.BadRequest("PayService.Notify", "Failed update trade record")
		}
		body := struct {
			OutPayNo string `json:"out_pay_no"` //三方单号
			TotalFee int64  `json:"total_fee"`  //支付金额
			PayTime  string `json:"pay_time"`   //支付时间
			Raw      string `json:"raw"`        //原始数据
		}{
			OutPayNo: tm.OutPayNo,
			TotalFee: tm.TotalFee,
			PayTime:  utils.FormatDate(tm.PayAt),
			Raw:      raw,
		}
		bodyByte, _ := json.Marshal(body)
		if err != nil {
			rsp.Result = wxpay.NotifyReplyFail("转码交易数据时发生错误")
			return errors.BadRequest("PayService.Notify", "Failed json.marshal trade record")
		}
		//进行真实回调任务
		ev := &pay_srv.NotifyEvent{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Name:      "pay_notify",
			PayNo:     tm.PayNo,
			Url:       tm.NotifyURL,
			Body:      string(bodyByte),
		}
		err = queue.Publish(ctx, "pay_notify", ev)
		if err != nil {
			log.Fatal(err)
		}
		//返回支付成功信息
		rsp.Result = wxpay.NotifyReplySuccess()
		rsp.OutPayNo = tm.OutPayNo
		// } else if payConf.Provider == TradeAlipay { //阿里支付
	} else {
		return errors.BadRequest("PayService.Create", "pay channel is undefined")
	}
	return nil
}

//Query ..
func (s *PayService) Query(ctx context.Context, req *pay_srv.QueryReq, rsp *pay_srv.PayRsp) error {
	// log.Log("[access] PayService.Query")
	// req.GetUniTradeNo()
	//查找订单
	tm := &model.TradeModel{}
	err := orm.Model(tm).Where("pay_no", "554988925916024832").Find()
	if err != nil {
		return errors.BadRequest("PayService.Query", fmt.Sprintf("not found trade record, pay_no=%s, %s", "554988925916024832", err.Error()))
	}
	fmt.Println(tm)
	return nil
}
