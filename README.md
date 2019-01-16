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
            micro call go.micro.user_srv UserService.Login '{"platform":"phone","login":"18310497688","password":"098f6bcd4621d373cade4e832627b4f6"}'
        第三方登陆
    用户注册
        手机号注册
        第三方注册
    用户资料
        micro call go.micro.user_srv UserService.User '{"jwt":"eyJhbGciOiJIUzI1NiIsImtpZCI6IjcyNDUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJ1c2VyLXNydiIsInN1YiI6IjEiLCJleHAiOjE1NDQwMjYyNDUsIm5iZiI6MTU0MzQyMzI0NSwiaWF0IjoxNTQzNDIxNDQ1fQ.uVIPNw-JsTgU1yCbkHjQmRsPdkdP7kUG7jLkq5TYXe4"}'
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