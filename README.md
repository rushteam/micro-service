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