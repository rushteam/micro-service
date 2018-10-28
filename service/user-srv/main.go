package main

import (
	"log"

	"gitee.com/coremicro/auth/common/db"
	"gitee.com/rushteam/micro-service/common/pb/user_srv"
	"gitee.com/rushteam/micro-service/service/user-srv/handler"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.user_srv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.Flags(
			cli.StringFlag{
				Name:   "config_path",
				EnvVar: "CONFIG_PATH",
				Usage:  "The config PATH e.g ./config.yaml",
			},
		),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			// var configFile = "./config.yaml"
			// if len(c.String("config_path")) > 0 {
			// 	configFile = c.String("config_path")
			// }
			user_srv.RegisterUserServiceHandler(service.Server(), new(handler.UserServiceHandler))
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))

			db.InitDB("root:123321@tcp(192.168.33.10:3306)/auth?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s")
			model.SetDB(db.DB)
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
