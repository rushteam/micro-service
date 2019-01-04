package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.web.auth_web"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

//PayNotifyHandler ..
func PayNotifyHandler(c *gin.Context) {
	// c.GetQuery()
	raw, err := c.GetRawData()
	if err != nil {
		return
	}
	// redirectURI := c.Query("redirect_uri")
	// if redirectURI == "" {
	// 	// c.AbortWithError(http.StatusBadRequest, errors.New("缺少参数 redirect_uri"))
	// 	c.String(http.StatusOK, "缺少参数 redirect_uri")
	// 	return
	// }
}
func main() {
	// Creates an application without any middleware by default.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/pay/notify/", PayNotifyHandler)
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
