package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	micro "github.com/micro/go-web"
	"github.com/mlboy/micro-service/service/pay-api/handler"
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
	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			r.Use(gin.Logger())
			r.Use(gin.Recovery())
			rpcHandler := &handler.RpcHandler{}
			r.POST("/rpc", rpcHandler.Create)

			payNotifyHandler := &handler.PayNotifyHandler{}
			//TODO: /pay/notify/wcpay/:channel 对channel的校验？
			// r.POST("/pay/notify/wcpay/:channel", payNotifyHandler.Wcpay)
			r.POST("/pay/notify/:channel", payNotifyHandler.Notify)

			// var services []interface{}
			// for i, srv := range services {
			// 	r.Handle(srv.Method, srv, Path, handler.RPC)
			// }
			// r.POST("/pay/order/create", handler.Create)
			// r.POST("/pay/order/query", handler.Query)
		}),
		//web
		micro.Handler(r),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
