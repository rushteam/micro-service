# micro-service
protoc --proto_path=./proto --micro_out=./common/pb/ --go_out=plugins=micro:./common/pb/ ./proto/user_srv/*.proto

micro call pay_center PayCenter.CreatePayOrder '{}'

micro call pay_center PayCenter.CreatePayOrder '{"client_id":"hoo","channel":"201","out_trade_no":"001","total_fee":1,"subject":"测试","from_ip":"127.0.0.1","open_id":"o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ"}'

micro call go.micro.user_srv UserService.Create '{"login_list":[{"platform":"phone"}],"userinfo":{"nickname":"小花猫"}}'

micro call go.micro.user_srv UserService.Bind '{"login_list":[{"platform":"phone"}],"userinfo":{"uid":1}}'

go run ./service/user-srv/main.go

http://127.0.0.1:9080/oauth2/authorize?response_type=code&client_id=1234&redirect_uri=http%3a%2f%2fwww.1thx.com

http://127.0.0.1:9080/oauth2/authorize?client_id=1234&response_type=code

http://127.0.0.1:9080/oauth2/token?client_id=1234&client_secret=test&grant_type=client_credentials&scope=token


用户服务 user-srv
    用户登陆
        手机号登陆
            micro --registry=consul call go.micro.srv.user_srv UserService.LoginByPassword '{"loginname":"18310497688","password":"test"}'
        第三方登陆
            micro --registry=consul call go.micro.srv.user_srv UserService.LoginByOAuth '{"openid":"test","password":"test","platform":"wx"}'

            micro --registry=consul call go.micro.srv.user_srv UserService.LoginByOAuth '{"appid":"test","sercet":"test","code":"test","platform":"wx"}'
    用户注册
        手机号注册
            micro  --registry=consul call go.micro.srv.user_srv UserService.Login '{"login_list":[{"platform":"phone","login":"18310497699","password":"test"}],"user":{"nickname":"测试"}}'
        第三方注册
    用户资料
        micro --registry=consul call go.micro.srv.user_srv UserService.User '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLXNydiIsInN1YiI6IjEiLCJhdWQiOiIiLCJleHAiOjE1NjkzNjYxMDksIm5iZiI6MTU2ODc2MTMwOSwiaWF0IjoxNTY4NzYxMzA5fQ.aw064Bo6yc-2UYNUK1cYlJTIc8eB6tZBjm1K16OxEoo"}'
    用户修改资料
    用户修改密码

权限服务 auth-srv
    oauth2鉴权
        app鉴权
        用户统一登陆
雪花算法序号生成 snowflake-srv
    生成号码
        parma 类别
订单服务 order-srv
    预算订单（下单页）
        micro call go.micro.srv.order_srv OrderService.Budget '{"order":{"items":[{"qty":2,"sku_id":1,"price":3000},{"qty":2,"sku_id":1,"price":3000}]}}'
    创建订单（确认按钮）
        micro call go.micro.srv.order_srv OrderService.Create '{"order":{"items":[{"qty":2,"sku_id":1,"price":3000},{"qty":2,"sku_id":1,"price":3000}]}}'
    订单列表
         micro call go.micro.srv.order_srv OrderService.Order '{"order_no":""}'
        
    订单详情

支付服务 pay-srv
    创建支付单
     micro call go.micro.srv.pay_srv PayService.Create '{"out_pay_no":"test_001","client_id":"hoo","access_token":"hoo","channel":"201","total_fee":100,"subject":"测试支付","trade_type":"JSAPI","notify_url":"https://1thx.com/","extra":"{\"openid\":\"o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ\"}"}'

    支付回调(支付宝/微信) - 回调具体服务
        micro call go.micro.srv.pay_srv PayService.Notify '{"client_id":"hoo","channel":"201","pvd_name":"wxpay","raw":"<xml><return_code>SUCCESS</return_code><result_code>SUCCESS</result_code><out_trade_no>1001</out_trade_no></xml>"}'
    
    支付单查询
        micro call go.micro.srv.pay_srv PayService.Query '{"client_id":"hoo","channel":"201","out_pay_no":"554988925916024832"}'

支付网关
    支付回调

启动服务

go run service/pay-srv/main.go --config_path="service/pay-srv/config.yaml"

mysql -uroot -h127.0.0.1  -e "CREATE DATABASE rushteam DEFAULT CHARSET utf8 DEFAULT COLLATE utf8_general_ci;"


127.0.0.1:9080/pay/order/create

{"service":"go.micro.srv.pay_srv","endpoint":"","method":"PayService.Create","request":{"out_pay_no":"test_001","client_id":"hoo","access_token":"hoo","channel":"201","total_fee":100,"subject":"测试支付","trade_type":"JSAPI","notify_url":"https://1thx.com/","extra":"{\"openid\":\"o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ\"}"}}


docker run --env="MICRO_REGISTRY_ADDRESS=127.0.0.1:8500" rushteam/micro-pay-srv:latest

172.17.0.1

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o micro-pay-srv  -gcflags '-N -l' ./service/pay-srv/*.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o micro-pay-api  -gcflags '-N -l' ./service/pay-api/*.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o micro-pay-notify-web  -gcflags '-N -l' ./service/pay-notify-web/*.go
test


https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6


使用--registry=consul



package handler
import (
	"context"
	"github.com/micro/go-log"
	example "{{.Dir}}/proto/example"
)
type Example struct{}
// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Log("Received Example.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}