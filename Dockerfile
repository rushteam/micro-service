FROM golang:1.11.2-alpine as builder

 #WORKDIR /go/src/micro-service

 #COPY ./ /go/src/micro-service

RUN apk add --no-cache gcc musl-dev 

WORKDIR /app/src/micro-service

COPY ./ /app/src/micro-service

# RUN set -ex && \
#     go build -v -o /go/bin/micro-service \
#     -gcflags '-N -l' \
#     ./*.go
RUN set -ex && \
    # go get -d -v && \
    go build -v -o /go/bin/micro-service \
    -mod=vendor \
    -gcflags '-N -l' \
    ./service/user-srv/main.go
