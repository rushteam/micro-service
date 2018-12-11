package model

//TradeModel ..
type TradeModel struct {
	// gorm.Model
	//交易单号
	TradeNo    string `gorm:"PRIMARY_KEY;"`
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
	//第三方通知地址
	NotifyURL string
	// 商户id
	ProviderMchId string
	//提供商公众号
	ProviderAppid string
	//pay_field
}

//TableName ..
func (PayModel) TableName() string {
	return "pay_trade"
}

//GetOrderByOrderNo ...
func (sess *Session) GetTradeByOrderNo(no string) (TradeModel, error) {
	trade := TradeModel{}
	return trade, nil
}
