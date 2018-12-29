package queue

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
)

var (
	//Queues ..
	Queues = make(map[string]micro.Publisher, 0)
)

//Register ..
func Register(name string, p micro.Publisher) {
	Queues[name] = p
}

//Publish ..
func Publish(ctx context.Context, name string, ev interface{}) error {
	if p, ok := Queues[name]; ok {
		p.Publish(ctx, ev)
	}
	return fmt.Errorf("[queue] 发送失败，未找到名为%s队列", name)
}

//Get ..
func Get(name string) micro.Publisher {
	if p, ok := Queues[name]; ok {
		return p
	}
	return nil
}
