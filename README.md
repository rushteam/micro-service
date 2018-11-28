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
        第三方登陆
    用户注册
        手机号注册
        第三方注册
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
    创建订单
    订单列表
    订单详情

支付服务 pay-srv
    创建支付单
    支付回调（暴露web post api）
        支付宝
        微信
    回调具体服务
    
