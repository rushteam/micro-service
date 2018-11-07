package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-api"
	micro "github.com/micro/go-micro"

	"github.com/micro/examples/template/web/handler"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.oauth_api"
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
			dbSource := "root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			pool := db.InitDb("mysql", dbSource, true)
			model.Init(pool)
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))
		}),
	)
	opt := api.WithEndpoint(&api.Endpoint{
		// The RPC method
		Name: "Greeter.Hello",
		// The HTTP paths. This can be a POSIX regex
		Path: []string{"/greeter"},
		// The HTTP Methods for this endpoint
		Method: []string{"GET", "POST"},
		// The API handler to use
		Handler: rpc.Handler,
	}
	user_srv.RegisterUserServiceHandler(service.Server(), new(handler.UserService),opt)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
