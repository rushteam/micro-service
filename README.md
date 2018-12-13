# micro-service
protoc --proto_path=./proto --micro_out=./common/pb/ --go_out=plugins=micro:./common/pb/ ./proto/user_srv/*.proto


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
        micro call go.micro.api.order_srv OrderService.Budget '{"order":{"items":[{"qty":2,"sku_id":1,"price":3000},{"qty":2,"sku_id":1,"price":3000}]}}'
    创建订单（确认按钮）
        micro call go.micro.api.order_srv OrderService.Create '{"order":{"items":[{"qty":2,"sku_id":1,"price":3000},{"qty":2,"sku_id":1,"price":3000}]}}'
    订单详情
        micro call go.micro.api.order_srv OrderService.Order ''
    订单列表
        micro call go.micro.api.order_srv OrderService.OrderList ''

支付服务 pay-srv
    文档设计参考 http://www.xxpay.org/dev/api.html#api-tab=tab-api

    统一下单 PayService.Create
       trade_type=JSAPI {"openId":"o2RvowBf7sOVJf8kJksUEMceaDqo"}
       trade_type=NATIVE {"productId":"120989823"}

       micro call go.micro.api.pay_srv PayService.Create '{"out_trade_no":"test_001","client_id":"hoo","access_token":"hoo","channel":"201","total_fee":100,"subject":"测试支付","trade_type":"JSAPI","extra":"{\"openid\":\"o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ\"}"}'

    支付回调 PayService.Notify

        回调具体服务
        micro call go.micro.api.pay_srv PayService.Notify '{}'


    支付单查询 PayService.Query

    支付退款 PayService.Refund

支付网关 pay-gateway
    域名 pay.xixihi.com
    路由
        支付宝 支付回调 /pay/notify/alipay
        微信 支付回调 /pay/notify/wx
            调用 pay-srv PayService.Notify


//统一下单返回数据
{
	"provider_name": "wxpay",
	"channel": "201",
	"pay_field": {
		"app_id": "wx5e596d33cb663cd1",
		"out_trade_no": "test_002",
		"total_fee": 100,
		"field_str": "{\"appId\":\"wx5e596d33cb663cd1\",\"timeStamp\":1544616541,\"nonceStr\":\"HRQWW8qGrVcg6Gz2Ca0YhaqvLk23jE10\",\"package\":\"prepay_id=wx1220090191776662328695910547854915\",\"signType\":\"MD5\",\"paySign\":\"488763BC707E72EBFDB7A8246A67AD36\"}"
	}
}
