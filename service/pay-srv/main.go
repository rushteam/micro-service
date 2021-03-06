package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rushteam/micro-service/common/pb/pay_srv"
	"github.com/rushteam/micro-service/service/pay-srv/config"
	"github.com/rushteam/micro-service/service/pay-srv/queue"

	// "github.com/rushteam/micro-service/common/db"
	"github.com/rushteam/micro-service/common/micro/wrap"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/rushteam/micro-service/service/pay-srv/handler"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.srv.pay_srv"
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
			// cli.StringFlag{
			// 	Name:   "app_db",
			// 	EnvVar: "MS_PAY_SRV_DB",
			// 	Usage:  "Db config for mysql e.g username:password@tcp(host:port)/database",
			// 	Value:  "root:dream@tcp(127.0.0.1:3306)/rushteam",
			// },
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
			configFile := c.String("config_path")
			err := config.App.Load(configFile)
			if err != nil {
				log.Fatal(err)
			}
			dbConf, err := config.App.Db.Default()
			if err != nil {
				log.Fatal(err)
			}
			queue.RegisterPublisher("pay_notify", micro.NewPublisher("go.micro.evt.pay_srv.pay_notify", service.Client()))
			micro.RegisterSubscriber("go.micro.evt.pay_srv.pay_notify", service.Server(), new(queue.Consumer))
			pay_srv.RegisterPayServiceHandler(service.Server(), &handler.PayService{Service: service})
			// fmt.Println(uuid.Parse(service.Server().Options().Id).Version())
			// fmt.Printf("%s", c.String("server_id"))
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
