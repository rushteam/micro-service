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
	c.Data(200, "application/xml; charset=utf-8", []byte(rst.Result))
	// c.String(200, "%s", rst.Result)
}
func main() {
	// Creates an application without any middleware by default.
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	payNotifyHandler := &PayNotifyHandler{}
	//TODO: /pay/notify/wcpay/:channel 对channel的处理
	r.POST("/pay/notify/wcpay", payNotifyHandler.Wcpay)
	// r.POST("/pay/notify/alipay", PayNotifyHandler)
	// r.HandleFunc("/objects/{object}", objectHandler)
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
			// var configFile = "./config.yaml"
			// if len(c.String("config_path")) > 0 {
			// 	configFile = c.String("config_path")
			// }
			// dbSource := "root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			// pool := db.InitDb("mysql", dbSource, true)
			// model.Init(pool)
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))
		}),
		//web
		micro.Handler(r),
		micro.Address(":9080"),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
