package main

import (
	"log"

	"gitee.com/rushteam/micro-service/common/db"
	"gitee.com/rushteam/micro-service/service/user-srv/model"

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
				Name:   "db",
				EnvVar: "MS_USER_SRV_DB",
				Usage:  "db配置",
				Value: "root:dream@tcp(127.0.0.1:3306)/rushteam",
			},
		),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			dbConf := c.String("db")
			dbSource := dbConf + "?" + "parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			pool := db.InitDb("mysql",dbSource,true)
			model.Init(pool)
			user_srv.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
