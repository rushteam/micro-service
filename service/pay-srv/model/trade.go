package model

import (
	"gitee.com/rushteam/micro-service/common/db"
)

//TradeModel ..
type TradeModel struct {
	// gorm.Model
	//交易单号
	TradeNo string `gorm:"PRIMARY_KEY;"`
	//商户号
	ClientId string
	//外部交易号、订单号
	OutTradeNo string
	//交易渠道 (微信支\付宝)
	Channel string
	//支付金额
	TotalFee int64
	//交易方式 (JSAPI\NATIVE)
	TradeType string
	//交易标题
	Subject string
	//交易来源ip
	FromIp string
	//支付状态
	PayState int64
	//支付时间
	PayAt int64
	//提供商名字
	Provider string
	// Provider商户id
	PvdMchId string
	//提供商公众号
	PvdAppid string
	//订单号
	PvdOutTradeNo string
	//交易单号
	PvdTradeNo string
	//pay_field
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
func (m TradeModel) Save() db.Model {
	return db.Model
}
