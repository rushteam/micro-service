package main

import (
	"fmt"
	"log"
	"time"

	"gitee.com/rushteam/micro-service/service/pay-api/handler"

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

func main() {
	// Creates an application without any middleware by default.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	payNotifyHandler := &handler.PayNotifyHandler{}
	r.POST("/pay/notify/wcpay", payNotifyHandler.Wcpay)
	// r.POST("/pay/notify/alipay", PayNotifyHandler)
	// r.HandleFunc("/objects/{object}", objectHandler)
	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.Flags(
			cli.StringFlag{
				Name:   "port",
				EnvVar: "MS_HTTP_PORT",
				Usage:  "http port",
				Value:  ":8080",
			},
		),
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
			fmt.Println(c.String("port"))
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
		// micro.
		micro.Handler(r),
		micro.Address(":9080"),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
