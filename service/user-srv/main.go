package main

import (
	"encoding/base64"
	"io/ioutil"
	"time"

	"github.com/rushteam/gosql"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/auth/jwt"
	"github.com/micro/go-micro/v2/logger"
	"github.com/rushteam/micro-service/common/micro/wrap"
	"github.com/rushteam/micro-service/service/user-srv/handler"

	// "upper.io/db.v3/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	//ServiceName service's name
	ServiceName = "go.micro.srv.usersrv"
	//ServiceVersion service's version
	ServiceVersion = "latest"

	excludeMethods = []string{"UserService.Signin", "UserService.Signup", "UserService.OAuthAuthorize"}
)

func main() {
	privateKey, _ := ioutil.ReadFile("./rsa_private_key.pem")
	authd := jwt.NewAuth(
		auth.PrivateKey(base64.StdEncoding.EncodeToString(privateKey)),
		auth.Exclude(excludeMethods...),
	)
	service := micro.NewService(
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
		micro.Name(ServiceName),
		micro.Version(ServiceVersion),
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
			_ = gosql.NewCluster(
				gosql.AddDb("mysql", "root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"),
			)
			// defer sess.Close()
			handler.RegisterUserServiceHandler(service)
			return nil
		}),
	)
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

//dazzlego
//温馨家园 三栋 2单元
