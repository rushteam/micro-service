package handler

import (
	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

//PayNotifyHandler ..
type PayNotifyHandler struct{}

//Wcpay ..
func (h PayNotifyHandler) Wcpay(c *gin.Context) {
	// c.GetQuery()
	// author := c.GetHeader("Authorization") //Authorization: Signature xxx
	// author := c.GetHeader("X-Signature") //Authorization: Signature
	raw, err := c.GetRawData()
	if err != nil {
		c.String(500, "%s", err.Error())
		return
	}
	if len(raw) == 0 {
		c.String(500, "%s", "NO DATA")
		return
	}
	// fmt.Println(raw)
	paySrv := pay_srv.NewPayService("go.micro.srv.pay_srv", client.DefaultClient)
	rst, err := paySrv.Notify(c, &pay_srv.NotifyReq{})
	if err != nil {
		c.String(500, "%s", err.Error())
		return
	}
	// fmt.Println(rst.Result)
	c.String(200, "%s", rst.Result)
}
