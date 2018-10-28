package main

import (
	"context"
	"log"

	"gitee.com/rushteam/micro-service/common/pb/user_srv"
	"gitee.com/rushteam/micro-service/service/user-srv/handler"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "user-srv"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
		micro.Flags(
			cli.StringFlag{
				Name:   "config_path",
				EnvVar: "CONFIG_PATH",
				Usage:  "The config PATH e.g ./config.yaml",
			},
		),
	)
	var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			// var configFile = "./config.yaml"
			// if len(c.String("config_path")) > 0 {
			// 	configFile = c.String("config_path")
			// }
			user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
