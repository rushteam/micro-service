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
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handler := &handler.Handler{}
	// var services []interface{}
	// for i, srv := range services {
	// 	r.Handle(srv.Method, srv, Path, handler.RPC)
	// }
	r.POST("/pay/order/create", handler.Create)
	// r.POST("/pay/order/query", handler.Query)

	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			// address = c.String("address")
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
		// micro.Address(address),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
