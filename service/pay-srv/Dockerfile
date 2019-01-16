FROM golang:1.11-alpine as builder

 #WORKDIR /go/src/micro-service

 #COPY ./ /go/src/micro-service


RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache gcc musl-dev 

WORKDIR /app/src/micro-service

COPY ./ /app/src/micro-service

RUN set -ex
# RUN go build -v -o /go/bin/micro-user-srv -mod=vendor -gcflags '-N -l' ./service/user-srv/main.go
RUN go build -v -o /go/bin/pay-srv -mod=vendor -gcflags '-N -l' ./service/pay-srv/main.go

FROM golang:1.11-alpine

COPY --from=builder /go/bin/ /usr/bin/
RUN mkdir -p /data/pay-srv/
COPY --from=builder /app/src/micro-service/service/pay-srv/*.yaml /data/pay-srv/

CMD ["pay-srv","--config_path=/data/pay-srv/config.yaml"]