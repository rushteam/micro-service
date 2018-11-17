FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-docker

COPY ./ /go/src/component-docker

RUN set -ex && \
go build -v -o /go/bin/component-docker \
-gcflags '-N -l' \
./*.go
