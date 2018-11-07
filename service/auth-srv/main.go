package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	micro "github.com/micro/go-micro"
	
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.web.auth_srv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

func main() {
	// Creates an application without any middleware by default.
	r := iris.New()
	 // Recover middleware recovers from any panics and writes a 500 if there was one.
	r.Use(recover.New())

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/objects/{object}", objectHandler)

	service := web.NewService(
		web.Handler(r)
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
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
	
	
	service.HandleFunc("/foo", fooHandler)

	// app := iris.Default()
    // app.Get("/ping", func(ctx iris.Context) {
    //     ctx.JSON(iris.Map{
    //         "message": "pong",
    //     })
    // })
    // // listen and serve on http://0.0.0.0:8080.
	// app.Run(iris.Addr(":8080"))
	
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
