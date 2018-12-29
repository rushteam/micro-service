package wrap

import (
	"context"
	"time"

	log "github.com/micro/go-log"
	"github.com/micro/go-micro/server"
)

//Access is a middleware to log request/responses.
func Access(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		log.Logf("[access] %s start:%s", req.Method(), start)
		res := fn(ctx, req, rsp)
		if res != nil {
			log.Logf("[service] %s %s", req.Method(), res.Error())
		}
		elapsed := time.Since(start).Round(time.Millisecond).String()
		log.Logf("[elapsed] %s elapsed:%s", req.Method(), elapsed)
		return res
	}
}
