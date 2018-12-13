package mch

import "encoding/xml"

//Notify ..
type Notify struct {
	XMLName struct{} `xml:"xml"`
	//必有
	ReturnCode string `xml:"return_code"` //返回状态码 SUCCESS/FAIL
	ReturnMsg  string `xml:"return_msg"`  //返回信息
	//以下字段在return_code为SUCCESS的时候有返回
	AppID      string `xml:"appid"`        //公众账号ID
	MchID      string `xml:"mch_id"`       //商户号
	DeviceInfo string `xml:"device_info"`  //设备号
	NonceStr   string `xml:"nonce_str"`    //随机字符串
	Sign       string `xml:"sign"`         //签名 HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	SignType       string `xml:"sign_type"`         //签名类型

	ResultCode string `xml:"result_code"`  //业务结果 SUCCESS	SUCCESS/FAIL
	ErrCode    string `xml:"err_code"`     //错误代码 当result_code为FAIL时返回错误代码，详细参见错误列表
	ErrCodeDes string `xml:"err_code_des"` //错误代码描述
	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Openid     string `xml:"openid"`	//微信openid 用户在商户appid下的唯一标识
	IsSubscribe string `xml:"is_subscribe"`	//是否关注公众账号 Y	用户是否关注公众账号，Y-关注，N-未关注
	TradeType string `xml:"trade_type"` //交易类型 JSAPI 公众号支付 NATIVE 扫码支付 APP APP支付
	BankType string `xml:"bank_type"` //付款银行 CMC	银行类型，采用字符串类型的银行标识，银行类型见银行列表
	TotalFee int64 `xml:"total_fee"` //订单金额	total_fee	是	Int	100	订单总金额，单位为分
	SettlementTotalFee int64 `xml:"settlement_total_fee"` //应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType string `xml:"fee_type"` //货币种类 CNY	货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee int64 `xml:"cash_fee"` //现金支付金额 订单现金支付金额，详见支付金额
	CashFeeType string `xml:"cash_fee_type"` //现金支付货币类型 CNY	货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CouponFee int64 `xml:"coupon_fee"` //总代金券金额	代金券金额<=订单金额，订单金额-代金券金额=现金支付金额，详见支付金额
	CouponCount int64 `xml:"coupon_count"` //代金券使用数量

	CouponType0 int64 `xml:"coupon_type_0"` //代金券类型	coupon_type_$n	否	String	CASH (CASH--充值代金券,NO_CASH---非充值代金券)
	CouponId0	string `xml:"coupon_id_0"`	//代金券ID	coupon_id_$n	否	String(20)	10000	代金券ID,$n为下标，从0开始编号
	CouponFee0	string `xml:"coupon_fee_0"`	//单个代金券支付金额	coupon_fee_$n	否	Int	100	单个代金券支付金额,$n为下标，从0开始编号

	TransactionId	string `xml:"transaction_id"`	//微信支付订单号 1217752501201407033233368018	微信支付订单号
	OutTradeNo	string `xml:"out_trade_no"`	//商户订单号
	Attach	string `xml:"attach"`	//商家数据包 (128)
	TimeEnd	string `xml:"time_end"`	//支付完成时间 格式为yyyyMMddHHmmss
}
//UnmarshalNotify ..
func UnmarshalNotify(raw string) (*Notify,error){
	rawBytes := []byte(raw)
	notify := &Notify{}
	err := xml.Unmarshal(rawBytes,notify)
	return notify,err
}