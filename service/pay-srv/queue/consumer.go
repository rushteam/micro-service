package queue

import (
	"context"
	"fmt"

	log "github.com/micro/go-log"

	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	// "github.com/go-log/log"
)

//Consumer 消费者
type Consumer struct{}

//Process  Method can be of any name
func (s *Consumer) Process(ctx context.Context, event *pay_srv.NotifyEvent) error {
	// md, _ := metadata.FromContext(ctx)
	log.Logf("recvied data: %+v\r\n", event)
	// utils.PostURL(url, params)

	return fmt.Errorf("xxx")
	// return nil
}
