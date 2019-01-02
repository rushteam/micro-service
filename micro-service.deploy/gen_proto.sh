#!/bin/bash
protoc --proto_path=./proto --micro_out=./common/pb/ --go_out=plugins=micro:./common/pb/ ./proto/user_srv/*.proto
