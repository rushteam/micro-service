package wxpay

import (
	"encoding/xml"
	"fmt"
)

// const (
// 	//SUCCESS 成功
// 	SUCCESS = "SUCCESS"
// 	//FAIL 失败
// 	FAIL = "FAIL"
// )

// var mchURL = "https://api.mch.weixin.qq.com"

//UnifiedOrder ..
type UnifiedOrder struct {
	XMLName xml.Name `xml:"xml" url:"-"`
	/*必要*/
	AppID          string `xml:"appid,omitempty" url:"appid,omitempty"`                       //公众账号ID 微信支付分配的公众账号ID（企业号corpid即为此appId）
	MchID          string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                     //商户号 微信支付分配的商户号
	NonceStr       string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`               //随机字符串 长度要求在32位以内。推荐随机数生成算法
	Sign           string `xml:"sign,omitempty" url:"-"`                                      //签名 通过签名算法计算得出的签名值，详见签名生成算法
	Body           string `xml:"body,omitempty" url:"body,omitempty"`                         //商品描述 128
	OutTradeNo     string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`         //商户订单号 要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。详见商户订单号
	FeeType        string `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                 //标价币种 	CNY	符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
	TotalFee       int64  `xml:"total_fee,omitempty" url:"total_fee,omitempty"`               //标价金额 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string `xml:"spbill_create_ip,omitempty" url:"spbill_create_ip,omitempty"` //终端IP APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	NotifyURL      string `xml:"notify_url,omitempty" url:"notify_url,omitempty"`             //通知地址 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	TradeType      string `xml:"trade_type,omitempty" url:"trade_type,omitempty"`             //交易类型 JSAPI 公众号支付 NATIVE 扫码支付 APP APP支付
	/*可选*/
	DeviceInfo string `xml:"device_info,omitempty" url:"device_info,omitempty"` //设备号 	自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	SignType   string `xml:"sign_type,omitempty" url:"sign_type,omitempty"`     //签名类型 默认为MD5，支持HMAC-SHA256和MD5。
	Detail     string `xml:"detail,omitempty" url:"detail,omitempty"`           //商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”
	Attach     string `xml:"attach,omitempty" url:"attach,omitempty"`           //附加数据 在查询API和支付通知中原样返回，可作为自定义参数使用。
	TimeStart  string `xml:"time_start,omitempty" url:"time_start,omitempty"`   //交易起始时间 格式为yyyyMMddHHmmss
	TimeExpire string `xml:"time_expire,omitempty" url:"time_expire,omitempty"` //交易结束时间 订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id。其他详见时间规则，建议：最短失效时间间隔大于1分钟
	GoodsTag   string `xml:"goods_tag,omitempty" url:"goods_tag,omitempty"`     //订单优惠标记 使用代金券或立减优惠功能时需要的参数，说明详见代金券或立减优惠
	ProductID  string `xml:"product_id,omitempty" url:"product_id,omitempty"`   //商品ID 32字符内 trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	LimitPay   string `xml:"limit_pay,omitempty" url:"limit_pay,omitempty"`     //指定支付方式 no_credit	上传此参数no_credit--可限制用户不能使用信用卡支付
	OpenID     string `xml:"openid,omitempty" url:"openid,omitempty"`           //用户标识 trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
	SceneInfo  string `xml:"scene_info,omitempty" url:"scene_info,omitempty"`   //场景信息 该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }}
}

//URL ..
func (o UnifiedOrder) URL() string {
	return fmt.Sprintf("%s/pay/unifiedorder", mchURL)
}

//UnifiedOrderRsp ..
type UnifiedOrderRsp struct {
	XMLName struct{} `xml:"xml"`
	//必有
	ReturnCode string `xml:"return_code"` //返回状态码 SUCCESS/FAIL
	ReturnMsg  string `xml:"return_msg"`  //返回信息
	//以下字段在return_code为SUCCESS的时候有返回
	AppID      string `xml:"appid"`        //公众账号ID 调用接口提交的公众账号ID
	MchID      string `xml:"mch_id"`       //商户号 调用接口提交的商户号
	DeviceInfo string `xml:"device_info"`  //设备号
	NonceStr   string `xml:"nonce_str"`    //随机字符串
	Sign       string `xml:"sign"`         //签名
	ResultCode string `xml:"result_code"`  //业务结果 SUCCESS	SUCCESS/FAIL
	ErrCode    string `xml:"err_code"`     //错误代码 当result_code为FAIL时返回错误代码，详细参见错误列表
	ErrCodeDes string `xml:"err_code_des"` //错误代码描述
	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeTypes string `xml:"trade_type"` //交易类型 JSAPI 公众号支付 NATIVE 扫码支付 APP APP支付
	PrepayID   string `xml:"prepay_id"`  //预支付交易会话标识 微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
	CodeURL    string `xml:"code_url"`   //二维码链接 URl：weixin：//wxpay/s/An4baqw	trade_type为NATIVE时有返回，用于生成二维码，展示给用户进行扫码支付
	MwebURL    string `xml:"mweb_url"`   //拉起微信支付收银台的中间页面，可通过访问该url来拉起微信客户端，完成支付,mweb_url的有效期为5分钟
}

func (r *UnifiedOrderRsp) Error() error {
	if r.ReturnCode == Success && r.ResultCode == Success {
		return nil
	}
	if r.ReturnCode == Fail {
		return fmt.Errorf(r.ReturnMsg)
	}
	if r.ResultCode == Fail {
		return fmt.Errorf("%s: %s", r.ErrCode, r.ErrCodeDes)
	}
	return fmt.Errorf("unknow error")
}

//Payment ...
type Payment struct {
	AppID     string `json:"appId" url:"appId,omitempty"`         //公众号名称，由商户传入
	TimeStamp int64  `json:"timeStamp" url:"timeStamp,omitempty"` //时间戳，自1970年以来的秒数
	NonceStr  string `json:"nonceStr" url:"nonceStr,omitempty"`   //随机串
	Package   string `json:"package" url:"package,omitempty"`     //app:Sign=WXPay,公众号:prepay_id=xxx
	SignType  string `json:"signType" url:"signType,omitempty"`   //微信签名方式：
	PaySign   string `json:"paySign" url:"paySign,omitempty"`     //微信签名
}

// //MakeSign ...
// func (o *PayConfigJs) MakeSign(APIKey string) string {
// 	var params map[string]string
// 	params = utils.Struct2Map(*o, "json")
// 	o.PaySign = utils.Sign(params, "", fmt.Sprintf("&key=%s", APIKey), md5.New())
// 	o.SignType = "MD5"
// 	return o.PaySign
// }
