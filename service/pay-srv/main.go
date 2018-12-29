package main

import (
	"gitee.com/rushteam/micro-service/service/pay-srv/queue"
	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	"gitee.com/rushteam/micro-service/service/pay-srv/config"
	"log"

	"gitee.com/rushteam/micro-service/common/db"
	"gitee.com/rushteam/micro-service/common/micro/wrap"
	"gitee.com/rushteam/micro-service/service/pay-srv/model"

	"gitee.com/rushteam/micro-service/service/pay-srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/registry"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.api.pay_srv"
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
				EnvVar: "MS_PAY_SRV_DB",
				Usage:  "Db config for mysql e.g username:password@tcp(host:port)/database",
				Value: "root:dream@tcp(127.0.0.1:3306)/rushteam",
				// Value: "root:dream@tcp(mysql:3306)/rushteam",
			},
			cli.StringFlag{
				Name:   "config_path",
				EnvVar: "CONFIG_PATH",
				Usage:  "The config PATH e.g ../config/config.yaml",
				Value: "./config.yaml",
				// Value: "root:dream@tcp(mysql:3306)/rushteam",
			},
		),
		micro.WrapHandler(wrap.Access),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			// dbConf := c.String("app_db")
			// dbSource := dbConf + "?" + "parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			// pool := db.InitDb("mysql",dbSource,true)
			// model.Init(pool)

			configFile := c.String("config_path")
			err := config.App.Load(configFile)
			if err != nil {
				log.Fatal(err)
			}
			//初始化db
			db.Init(config.App.DbConfig)
			queue.Register("pay_notify",micro.NewPublisher("pay_notify", service.Client()))
			pay_srv.RegisterPayServiceHandler(service.Server(), &handler.PayService{Service:service})
			//fmt.Printf("%s",c.String("server_id"))
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}