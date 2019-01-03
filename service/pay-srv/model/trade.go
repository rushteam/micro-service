package model

import (
	"time"

	"github.com/mlboy/godb/orm"
)

/*
CREATE TABLE `pay_trade` (
  `pay_id` bigint(20) NOT NULL COMMENT '交易流水号(三方单号)',
  `out_pay_no` varchar(32) NOT NULL COMMENT '外部订单号',
  `client_id` varchar(128) NOT NULL COMMENT '商户',
  `channel` varchar(32) NOT NULL COMMENT '支付渠道',
  `total_fee` int(11) NOT NULL DEFAULT 0 COMMENT '支付金额/分',
  `fee_type` char(16) NOT NULL DEFAULT 'RMB' COMMENT '币种',
  `trade_type` varchar(16) NOT NULL COMMENT '交易方式',
  `subject` varchar(256) NOT NULL COMMENT '标题',
  `from_ip` varchar(256) NOT NULL COMMENT '来源ip',
  `provider` varchar(16) NOT NULL COMMENT '支付商@wechat@alipay',
  `pvd_mch_id` varchar(64) NOT NULL COMMENT '商户',
  `pvd_app_id` varchar(64) NOT NULL COMMENT '公众号',
  `pvd_trade_no` varchar(64) NOT NULL COMMENT '三方交易单号',
  `pay_state` int(11) NOT NULL DEFAULT 0 COMMENT '支付状态@1支付成功@2支付失败',
  `pay_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付时间',
  `notify_url` varchar(256) NOT NULL COMMENT '通知地址',
  `notify_num` int(11) NOT NULL DEFAULT 0 COMMENT '通知次数',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '修改时间',
  UNIQUE KEY `out_pay_no` (`out_pay_no`),
  UNIQUE KEY `pvd_trade_no` (`pvd_trade_no`),
  KEY `pvd_mch_id` (`pvd_mch_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/
//TradeModel ..
type TradeModel struct {
	//交易单号
	PayID string `db:"pay_id"`
	//外部交易号、订单号
	OutPayNo string `db:"out_pay_no"`
	//商户号
	ClientID string `db:"client_id"`
	//交易渠道 (微信支\付宝)
	Channel string `db:"channel"`
	//支付金额
	TotalFee int64 `db:"total_fee"`
	//币种
	FeeType string `db:"fee_type"`
	//交易方式 (JSAPI\NATIVE)
	TradeType string `db:"trade_type"`
	//交易标题
	Subject string `db:"subject"`
	//交易来源ip
	FromIP string `db:"from_ip"`
	//支付状态
	PayState int64 `db:"pay_state"`
	//支付时间
	PayAt time.Time `db:"pay_at"`
	//通知url
	NotifyURL string `db:"notify_url"`
	//通知次数
	NotifyNum int `db:"notify_num"`
	//三方提供商名字
	Provider string `db:"provider"`
	//三方商户
	PvdMchID string `db:"pvd_mch_id"`
	//三方appid 公众号
	PvdAppid string `db:"pvd_app_id"`
	//三方订单号
	PvdOutTradeNo string `db:"pvd_out_trade_no"`
	//三方交易单号
	PvdTradeNo string `db:"pvd_trade_no"`
	//三方异步通知url
	PvdNotifyURL string `db:"-"`
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
