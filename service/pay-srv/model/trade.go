package model

import (
	"time"

	"github.com/mlboy/godb/orm"
)

//TradeModel ..
type TradeModel struct {
	//交易单号
	TradeNo string `db:"trade_no,pk"`
	//商户号
	ClientId string `db:"client_id"`
	//外部交易号、订单号
	OutTradeNo string `db:"out_trade_no"`
	//交易渠道 (微信支\付宝)
	Channel string `db:"channel"`
	//支付金额
	TotalFee int64 `db:"total_fee"`
	//交易方式 (JSAPI\NATIVE)
	TradeType string `db:"trade_type"`
	//交易标题
	Subject string `db:"subject"`
	//交易来源ip
	FromIp string `db:"from_ip"`
	//支付状态
	PayState int64 `db:"pay_statte"`
	//支付时间
	PayAt time.Time `db:"pay_at"`
	//通知url
	NotifyUrl string `db:"notify_url"`
	//通知次数
	NotifyNum int `db:"notify_num"`
	//三方提供商名字
	Provider string `db:"provider"`
	//三方商户
	PvdMchId string `db:"pvd_mch_id"`
	//三方appid 公众号
	PvdAppid string `db:"pvd_app_id"`
	//三方订单号
	PvdOutTradeNo string `db:"pvd_out_trade_no"`
	//三方交易单号
	PvdTradeNo string `db:"pvd_out_trade_no"`
	//三方异步通知url
	PvdNotifyUrl string `db:"pvd_notify_url"`
}

//TableName ..
func (TradeModel) TableName() string {
	return "pay_trade"
}

//GetTradeByOrderNo ...
func (m *TradeModel) GetTradeByOrderNo(no string) (*TradeModel, error) {
	return m, nil
}

//Save ..
func (m *TradeModel) Save() error {
	_, err := orm.Model(m).Update()
	if err != nil {
		return err
	}
	return nil
}
