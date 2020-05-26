package main

import (
	"fmt"
	"log"

	"github.com/rushteam/micro-service/common/micro/wrap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/rushteam/micro-service/service/order-srv/handler"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.srv.order_srv"
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
				Value:  "",
				// Value: "root:dream@tcp(mysql:3306)/rushteam",
			},
		),
		micro.WrapHandler(wrap.Access),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			//service.Server().Options().Id
			//service.Server().Options().Name
			// fmt.Println(service.Server().Options().Id)
			// fmt.Println(srvs)
			// fmt.Println(srvs[0].Nodes[0].Id)
			dbSource := c.String("app_db")
			if dbSource == "" {
				dbSource = "root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			}
			order_srv.RegisterOrderServiceHandler(service.Server(), &handler.OrderService{Service: service})
			fmt.Printf("%s", c.String("server_id"))
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
