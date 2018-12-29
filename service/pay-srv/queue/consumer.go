package queue

import (
	"context"

	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	// "github.com/go-log/log"
	log "github.com/micro/go-log"
)

//Consumer 消费者
type Consumer struct{}

//Process  Method can be of any name
func (s *Consumer) Process(ctx context.Context, event *pay_srv.NotifyEvent) error {
	log.Logf("recvied data: %+v\r\n", event)
	// md, _ := metadata.FromContext(ctx)
	// log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
// func subEv(ctx context.Context, event *proto.Event) error {
// 	md, _ := metadata.FromContext(ctx)
// 	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
// 	// do something with event
// 	return nil
// }
