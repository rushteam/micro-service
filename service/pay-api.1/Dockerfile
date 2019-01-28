FROM golang:1.11-alpine as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache gcc musl-dev 

WORKDIR /app/src/micro-service
COPY ./ /app/src/micro-service
RUN set -ex
RUN go build -v -o /go/bin/pay-api -mod=vendor -gcflags '-N -l' ./service/pay-api/main.go


FROM golang:1.11-alpine

COPY --from=builder /go/bin/ /usr/bin/
CMD ["pay-api"]