package queue

import (
	"github.com/micro/go-micro/client"

	micro "github.com/micro/go-micro"
)

const payNotifyQueueKey = "pay_notify"

var (
	//PayNotify ...
	PayNotify micro.Publisher
)

//InitPayNotify ..
func InitPayNotify(c client.Client) micro.Publisher {
	PayNotify := micro.NewPublisher(payNotifyQueueKey, c)
	return PayNotify
}
