package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/mlboy/micro-service/common/micro/wrap"
	// "github.com/mlboy/micro-service/service/pay-srv/config"
	"github.com/micro/go-micro/config"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.srv.user_srv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

func main() {
	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.Flags(
			cli.StringFlag{
				Name:   "config_path",
				EnvVar: "CONFIG_PATH",
				Usage:  "The config PATH e.g ../config/config.yaml",
				Value:  "./config.yaml",
			},
		),
		micro.WrapHandler(wrap.Access),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			conf := config.NewConfig()
			config.LoadFile("./config.yaml")

			fmt.Printf("%v", conf.Get("db_configs"))
			// configFile := c.String("config_path")
			// err := config.App.Load(configFile)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// dbConf, err := config.App.Db.Default()
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// db, err := sql.Open(dbConf.DbType, dbConf.Nodes[0])
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// orm.InitDefaultDb(db)

			// user_srv.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
