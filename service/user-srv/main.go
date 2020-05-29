package main

import (
	"io/ioutil"
	"time"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/auth/jwt"
	"github.com/micro/go-micro/v2/logger"
	"github.com/rushteam/gosql"
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
	//base64.StdEncoding.EncodeToString(
	privateKey, _ := ioutil.ReadFile("./key")
	publicKey, _ := ioutil.ReadFile("./key.pub")
	authd := jwt.NewAuth(
		auth.PrivateKey(string(privateKey)),
		auth.PublicKey(string(publicKey)),
		// auth.Exclude(excludeMethods...),
	)
	srv := micro.NewService(
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
	srv.Init(
		micro.Action(func(c *cli.Context) error {
			gosql.NewCollect(
				gosql.NewCluster(
					gosql.AddDb("mysql", "root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"),
				),
			)
			// defer sess.Close()
			handler.RegisterUserServiceHandler(srv)
			return nil
		}),
	)
	acc, err := srv.Options().Auth.Generate("test", auth.WithType("api"))
	if err != nil {
		logger.Fatal(err)
	}
	tok, err := srv.Options().Auth.Token(auth.WithCredentials(acc.ID, acc.Secret))
	if err != nil {
		logger.Fatal(err)
	}
	srv.Options().Auth.Init(auth.ClientToken(tok))
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

//dazzlego
//温馨家园 三栋 2单元
