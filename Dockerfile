FROM golang:1.11-alpine as builder

WORKDIR /go/src/micro-service

COPY ./ /go/src/micro-service

# RUN set -ex && \
#     go build -v -o /go/bin/micro-service \
#     -gcflags '-N -l' \
#     ./*.go
RUN set -ex && \
    # go get -d -v && \
    go build -v -o -getmode=vendor /go/bin/micro-service \
    -gcflags '-N -l' \
    ./service/user-srv/main.go