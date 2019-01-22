package main

import (
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
			handler := &handler.Handler{}
			r.POST("/", handler.Create)
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
