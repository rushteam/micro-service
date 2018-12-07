package wrap

import (
	"context"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/server"
)

//Access is a middleware to log request/responses.
func Access(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		log.Logf("[access] %s start:%s", req.Method(), start)
		res := fn(ctx, req, rsp)
		elapsed := time.Since(start).Round(time.Millisecond).String()
		log.Logf("[elapsed] %s elapsed:%s", req.Method(), elapsed)
		return res
	}
}
