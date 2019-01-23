package main

import (
	"log"
	"time"

	"github.com/micro/go-micro/client"

	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	micro "github.com/micro/go-web"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.web.pay_notify"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

//PayNotifyHandler ..
type PayNotifyHandler struct{}

//Wcpay ..
func (h PayNotifyHandler) Wcpay(c *gin.Context) {
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
func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		// micro.Flags(
		// 	cli.StringFlag{
		// 		Name:   "config_path",
		// 		EnvVar: "CONFIG_PATH",
		// 		Usage:  "The config PATH e.g ./config.yaml",
		// 	},
		// ),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			payNotifyHandler := &PayNotifyHandler{}
			//TODO: /pay/notify/wcpay/:channel 对channel的校验？
			// r.POST("/pay/notify/wcpay/:channel", payNotifyHandler.Wcpay)
			r.POST("/pay/notify/wcpay/201", payNotifyHandler.Wcpay)
		}),
		//web
		micro.Handler(r),
		micro.Address(":9080"),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
