package main

import (
	"fmt"
	"log"

	"gitee.com/rushteam/micro-service/common/db"
	"gitee.com/rushteam/micro-service/service/order-srv/model"

	"gitee.com/rushteam/micro-service/common/pb/order_srv"
	"gitee.com/rushteam/micro-service/service/order-srv/handler"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	// "github.com/micro/go-micro/registry"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.api.order_srv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.Flags(
			cli.StringFlag{
				Name:   "app_db",
				EnvVar: "MS_ORDER_SRV_DB",
				Usage:  "Db config for mysql",
				Value: "root:dream@tcp(127.0.0.1:3306)/rushteam",
				// Value: "root:dream@tcp(mysql:3306)/rushteam",
			},
		),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			dbConf := c.String("app_db")
			dbSource := dbConf + "?" + "parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			pool := db.InitDb("mysql",dbSource,true)
			model.Init(pool)
			order_srv.RegisterOrderServiceHandler(service.Server(), new(handler.OrderService))
			fmt.Printf("%s",c.String("server_id"))
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
