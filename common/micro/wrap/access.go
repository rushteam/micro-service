package wrap

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
)

//Access is a middleware to log request/responses.
func Access(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		logger.Infof("[access] %v start: %v", req.Method(), start)
		res := fn(ctx, req, rsp)
		if res != nil {
			logger.Infof("[service] %v %v", req.Method(), res.Error())
		}
		elapsed := time.Since(start).Round(time.Millisecond).String()
		logger.Infof("[elapsed] %v elapsed: %v", req.Method(), elapsed)
		return res
	}
}
