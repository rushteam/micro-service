package queue

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
)

var (
	//Queues ..
	Queues = make(map[string]micro.Publisher, 0)
)

//RegisterPublisher 注册一个micro.Publisher
func RegisterPublisher(name string, p micro.Publisher) {
	Queues[name] = p
}

//Publish ..
func Publish(ctx context.Context, name string, ev interface{}) error {
	if p, ok := Queues[name]; ok {
		if p == nil {
			return fmt.Errorf("[queue] not found p")
		}
		err := p.Publish(ctx, ev)
		if err != nil {
			return fmt.Errorf("[queue] send fail, micro.Publisher.Publish %s", err.Error())
		}
		return nil
	}
	return fmt.Errorf("[queue] send fail, not found registered queue %s", name)
}

//Get ..
func Get(name string) micro.Publisher {
	if p, ok := Queues[name]; ok {
		return p
	}
	return nil
}
