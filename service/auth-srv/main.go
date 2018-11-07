package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-log"
	micro "github.com/micro/go-web"

	// micro "github.com/micro/go-micro"
	"github.com/gin-gonic/gin"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.web.auth_srv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

func main() {
	// Creates an application without any middleware by default.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r := iris.New()
	//  // Recover middleware recovers from any panics and writes a 500 if there was one.
	// r.Use(recover.New())

	// r.HandleFunc("/oauth2/", indexHandler)
	// r.HandleFunc("/objects/{object}", objectHandler)

	service := micro.NewService(
		micro.Handler(r),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.Address(":9080"),
		// micro.Flags(
		// 	cli.StringFlag{
		// 		Name:   "config_path",
		// 		EnvVar: "CONFIG_PATH",
		// 		Usage:  "The config PATH e.g ./config.yaml",
		// 	},
		// ),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) {
			// var configFile = "./config.yaml"
			// if len(c.String("config_path")) > 0 {
			// 	configFile = c.String("config_path")
			// }
			// dbSource := "root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"
			// pool := db.InitDb("mysql", dbSource, true)
			// model.Init(pool)
			// user_srv.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler(ctx))
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
