package main

import (
	"context"
	"github.com/micro/go-micro/client"
	"gitee.com/rushteam/micro-service/common/pb/user_srv"
	"github.com/RangelReale/osin"
	"net/http"
	"net/url"

	"github.com/micro/cli"
	"github.com/micro/go-log"
	micro "github.com/micro/go-web"

	// micro "github.com/micro/go-micro"
	"github.com/gin-gonic/gin"

	"gitee.com/rushteam/micro-service/service/auth-srv/oauth2"
)

var (
	//SERVICE_NAME service's name
	SERVICE_NAME = "go.micro.web.auth_srv"
	//SERVICE_VERSION service's version
	SERVICE_VERSION = "latest"
)

/**
client_id	true	int	应用申请时分配的appid.
response_type	true	string	目前固定为 "code"
redirect_uri	true	string	授权回调地址，域名需与设置的回调域名中任意一个域名一致,需要做url encode处理
scope	true	string	申请scope权限所需参数，可一次申请多个scope权限，目前只有 user_info 这个scope
state	false	string	用于保持请求和回调的状态，在回调时，会在Query Parameter中回传该参数。可以用这个参数验证请求有效性。这个参数可用于防止跨站请求伪造（CSRF）攻击
*/

//AuthorizeHandler ..
func AuthorizeHandler(c *gin.Context) {
	redirectURI := c.Query("redirect_uri")
	clientID := c.Query("client_id")
	responseType := c.Query("response_type")
	scope := c.Query("scope")
	state := c.Query("state")
	if redirectURI == "" {
		// c.AbortWithError(http.StatusBadRequest, errors.New("缺少参数 redirect_uri"))
		c.String(http.StatusOK, "缺少参数 redirect_uri")
		return
	}
	if clientID == "" {
		// c.AbortWithError(http.StatusBadRequest, errors.New("缺少参数 client_id"))
		c.String(http.StatusOK, "缺少参数 redirect_uri")
		return
	}
	if responseType == "" {
		// c.AbortWithError(http.StatusBadRequest, errors.New("缺少参数 response_type"))
		c.String(http.StatusOK, "缺少参数 response_type")
		return
	}
	if scope == "" {
		// c.AbortWithError(http.StatusBadRequest, errors.New("缺少参数 scope"))
		c.String(http.StatusOK, "缺少参数 scope")
		return
	}
	//http://127.0.0.1:9080/oauth2/authorize?redirect_uri=http://www.baidu.com&client_id=1&response_type=code&scope=token
	u, _ := url.Parse(redirectURI)
	params := u.Query()
	//302到
	//?code=xxxxx&state=test
	code := "test"
	//params := url.Values{}
	params.Add("code", code)
	params.Add("state", state)
	//query := params.Encode()
	redirectURL := u.Scheme + "://" + u.Host + u.Path + "?" + params.Encode()
	c.Redirect(http.StatusFound, redirectURL)
}
func TokenHandler(c *gin.Context) {
	c.String(http.StatusOK, `<html><body><h1>Hello World</h1></body></html>`)
}

func main() {
	// Creates an application without any middleware by default.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("service/auth-srv/templates/*")

	var loginPageHandler oauth2.LoginPageHandler
	loginPageHandler = func(ar *osin.AuthorizeRequest, c *gin.Context) bool {
		//todo 检测自己是否已经登录，如果登录提示 是否授权
		//oauth2.HandleDefaultLoginPage(ar,c.Writer,c.Request)
		r := c.Request
		if r.Method == "GET" {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"actionUrl": "/oauth2/authorize?" + r.URL.RawQuery,
			})
			return false
		}
		r.ParseForm()
		if login,ok := c.GetPostForm("login"); !ok {

		}
		if pwd,ok := c.GetPostForm("password"); !ok {

		}
		ctx := context.TODO()
		loginRsp,err := user_srv.NewUserService("go.micro.user_srv",client.NewClient()).Login(ctx, &user_srv.LoginReq{})
		// req := client.NewRequest("go.micro.user_srv", "UserService.Login", &user_srv.LoginReq{
		// 	Login: login,
		// 	Password: pwd,
		// })
		// rsp := &user_srv.LoginRsp{}
		// // Call service
		// if err := client.Call(ctx, req, rsp); err != nil {
		// 	fmt.Println("call err: ", err, rsp)
		// 	return
		// }
		if !userLogin(login,pwd) {
			//返回状态码
			c.String(200,"登录失败")
			// ar.Authorized = false
			return false
		}
		ar.Authorized = true
		return true
	}
	auth := oauth2.New(oauth2.NewDefaultOsinServer())
	auth.InitRouter(r)
	auth.SetLoginPageHandler(loginPageHandler)
	//authServer(r)
	//r.GET("/oauths2/authorize", AuthorizeHandler)

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