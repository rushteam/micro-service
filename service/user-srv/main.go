package main

import (
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
	authd := jwt.NewAuth(
		auth.PrivateKey(`MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALT41nn2U3v6zm5v
/9FfbUqUxspk3Q29VqcuyjwhFuOTjdu/eAr3uhKrGNcGmhv9rynmxQXtk8sijlDm
E01rO+XYIPDaazhqqwosfXrp2Bo7t+0xFhTLZWUfVUWHNK2uXcVikoUhdtHNKcQk
QzTKGtbv3lZ5bjflOV0oJZgTqjMbAgMBAAECgYBl8Qo395b9bsGcGkD7ewrAiWAV
oI2Y8MAAOu42wtj25yZw08FWREevvmumrJRhEhz6uIDhnvuy4MtULNZQtQgutIGd
GPi29ZCpTUykdvAoskUX2aWVlINeFIwUdI4HhiNPYz1flDvVil/8H2QsjaOhOAGI
++Gr+IvUEZ9qLQpBiQJBAOtqgjzSnNAJtXshCSvRE9r55FyNFQHPfPptWnWwm5sK
huprztQJ4j4n+vssoJmD6AUtDa5OiaDNqLknOhvlzJcCQQDEy6sPMq92YSAdq8pK
oGPp/3JRuiYu1hFAmBv+YVzulmTR8ogNXnG88bja3YIuqRh8CnKpMh0UIyp7ob6j
YOodAkBkhYY3EneDHaIwgVq5Kv2fczTfkB54N3DWPftyZYcMHOKfFomqYM4KXGf+
+H9bDcf07df5pe9+ilKRPP7DCszhAkBm35zdRf6pIF6chBgsaKmyQeGtzWR2aVr9
bEZ99MGSyzWK9oCelHdacPXUG6UY5TYyaXfE8Lh4tWVY2ZWBbIEVAkEA08YA88j7
2zeM/qyIXyaGw9l899h8aLPrN+kIsF45oRmfdl9d3qUejxgk8bmpWRBucLA+XeLx
tMkmw/XoEI6vDQ==`),
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
