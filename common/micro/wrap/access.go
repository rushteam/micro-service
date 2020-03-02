package wrap

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
)

//Access is a middleware to log request/responses.
func Access(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		log.Tracef("[access] %s start:%s", req.Method(), start)
		res := fn(ctx, req, rsp)
		if res != nil {
			log.Tracef("[service] %s %s", req.Method(), res.Error())
		}
		elapsed := time.Since(start).Round(time.Millisecond).String()
		log.Tracef("[elapsed] %s elapsed:%s", req.Method(), elapsed)
		return res
	}
}
