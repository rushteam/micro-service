# micro-service
protoc --proto_path=./proto --micro_out=./common/pb/ --go_out=plugins=micro:./common/pb/ ./proto/user_srv/*.proto

micro call pay_center PayCenter.CreatePayOrder '{"client_id":"hoo","channel":"201","out_trade_no":"001","total_fee":1,"subject":"测试","from_ip":"127.0.0.1","open_id":"o8UFh1m1fS3QiuSZ5Ik3rYgt3vjQ"}'