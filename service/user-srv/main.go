package main

import (
	"log"
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/rushteam/micro-service/common/micro/wrap"
	"github.com/rushteam/micro-service/service/user-srv/handler"
	"upper.io/db.v3/mysql"
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
				Name: "config_path",
				// EnvVars: "CONFIG_PATH",
				Usage: "The config PATH e.g ../application.yml",
				Value: "./application.yml",
			},
		),
		micro.WrapHandler(wrap.Access),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) error {
			settings, _ := mysql.ParseURL("root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s")
			sess, err := mysql.Open(settings)
			if err != nil {
				log.Fatalf("db.Open(): %q\n", err)
			}
			// defer sess.Close()
			handler.RegisterUserServiceHandler(service, sess)
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserService())
			return nil
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//dazzlego
//温馨家园 三栋 2单元
