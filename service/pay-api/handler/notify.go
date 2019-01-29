package handler

import (
	"github.com/micro/go-micro/client"
	"github.com/mlboy/micro-service/common/pb/pay_srv"

	"github.com/gin-gonic/gin"
)

//PayNotifyHandler ..
type PayNotifyHandler struct{}

//Notify ..
func (h PayNotifyHandler) Notify(c *gin.Context) {
	// c.GetQuery()
	// author := c.GetHeader("Authorization") //Authorization: Signature xxx
	// author := c.GetHeader("X-Signature") //Authorization: Signature
	channel := c.Param("channel")
	raw, err := c.GetRawData()
	if err != nil {
		c.String(500, "ERROR: %s", err.Error())
		return
	}
	if len(raw) == 0 {
		c.String(500, "ERROR: %s", "NO DATA")
		return
	}
	// fmt.Println(raw)
	paySrv := pay_srv.NewPayService("go.micro.srv.pay_srv", client.DefaultClient)
	rst, err := paySrv.Notify(c, &pay_srv.NotifyReq{
		Channel: channel,
		Raw:     string(raw),
	})
	if err != nil {
		c.String(500, "ERROR: %s", err.Error())
		return
	}
	// fmt.Println(rst.Result)
	c.Data(200, "application/xml; charset=utf-8", []byte(rst.Result))
	// c.String(200, "%s", rst.Result)
}
