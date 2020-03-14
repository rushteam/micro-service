package main

import (
	"encoding/base64"
	"io/ioutil"
	"time"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	log "github.com/micro/go-micro/v2/util/log"
	"github.com/rushteam/micro-service/common/micro/auth/jwt"
	"github.com/rushteam/micro-service/common/micro/wrap"
	"github.com/rushteam/micro-service/service/user-srv/handler"

	// "upper.io/db.v3/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mlboy/godb/db"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.srv.usersrv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

func main() {
	log.SetLevel(log.LevelTrace)
	privateKey, _ := ioutil.ReadFile("/Users/maliang/Documents/hoonet/rushteam/micro-service/rsa_private_key.pem")
	authd := jwt.NewAuth(
		auth.PrivateKey(base64.StdEncoding.EncodeToString(privateKey)),
		auth.Exclude("UserService.Signin"),
	)
	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.Auth(authd), //是否开启校验
		micro.Flags(
			&cli.StringFlag{
				Name:    "config_path",
				EnvVars: []string{"CONFIG_PATH"},
				Usage:   "The config PATH e.g ../application.yml",
				Value:   "./application.yml",
			},
		),
		micro.WrapHandler(wrap.Access),
	)
	// var ctx = context.TODO()
	service.Init(
		micro.Action(func(c *cli.Context) error {
			var settings = make(map[string][]string, 0)
			settings["default"] = []string{
				"root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s",
			}
			_ = db.InitPool("mysql", settings)
			// defer sess.Close()
			handler.RegisterUserServiceHandler(service)
			return nil
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//dazzlego
//温馨家园 三栋 2单元
