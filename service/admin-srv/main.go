package main

import (
	"context"
	"log"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/rushteam/micro-service/common/pb/wx_token"
	"github.com/rushteam/micro-service/service/admin-srv/handler"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "admin-srv"
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
			wx_token.RegisterWxTokenHandler(service.Server(), handler.NewWxToeknHandler(ctx))
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
