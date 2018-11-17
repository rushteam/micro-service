FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/micro-service

COPY ./ /go/src/micro-service

# RUN set -ex && \
#     go build -v -o /go/bin/micro-service \
#     -gcflags '-N -l' \
#     ./*.go
RUN set -ex && \
    go get && \
    go build -v -o /go/bin/micro-service \
    -gcflags '-N -l' \
    ./service/user-srv/*.go