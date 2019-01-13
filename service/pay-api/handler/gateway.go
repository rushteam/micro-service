package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-log/log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"

	// "github.com/micro/micro/internal/helper"

	"github.com/gin-gonic/gin"
)

type rpcRequest struct {
	Service  string
	Endpoint string
	Method   string
	Address  string
	Request  interface{}
}

//Handler ..
type Handler struct{}

//Create ..
func (h Handler) Create(c *gin.Context) {
	if c.Request.Method != "POST" {
		log.Log("Method not must POST")
		return
	}
	ct := c.GetHeader("Content-Type")
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}
	defer c.Request.Body.Close()
	var service, endpoint, address string
	var request interface{}
	switch ct {
	case "application/json":
		var rpcReq rpcRequest
		d := json.NewDecoder(c.Request.Body)
		d.UseNumber()
		if err := d.Decode(&rpcReq); err != nil {
			// badRequest(err.Error())
			log.Log(err.Error())
			return
		}
		service = rpcReq.Service
		endpoint = rpcReq.Endpoint
		address = rpcReq.Address
		request = rpcReq.Request
		if len(endpoint) == 0 {
			endpoint = rpcReq.Method
		}
		// JSON as string
		if req, ok := rpcReq.Request.(string); ok {
			d := json.NewDecoder(strings.NewReader(req))
			d.UseNumber()
			if err := d.Decode(&request); err != nil {
				log.Log("error decoding request string: " + err.Error())
				// badRequest("error decoding request string: " + err.Error())
				return
			}
		}
	default:
		service = c.PostForm("service")
		endpoint = c.PostForm("endpoint")
		address = c.PostForm("address")
		if len(endpoint) == 0 {
			endpoint = c.PostForm("method")
		}
		d := json.NewDecoder(strings.NewReader(c.PostForm("request")))
		d.UseNumber()
		if err := d.Decode(&request); err != nil {
			log.Log("error decoding request string: " + err.Error())
			// badRequest("error decoding request string: " + err.Error())
			return
		}
	}
	if len(service) == 0 {
		log.Log("invalid service")
		// badRequest("invalid service")
		return
	}
	if len(endpoint) == 0 {
		log.Log("invalid endpoint")
		// badRequest("invalid endpoint")
		return
	}
	// create request/response
	var response json.RawMessage
	var err error
	req := (*cmd.DefaultOptions().Client).NewRequest(service, endpoint, request, client.WithContentType("application/json"))
	// create context
	ctx := RequestToContext(c.Request)
	var opts []client.CallOption
	timeout, _ := strconv.Atoi(c.GetHeader("Timeout"))
	// set timeout
	if timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(timeout)*time.Second))
	}
	// remote call
	if len(address) > 0 {
		opts = append(opts, client.WithAddress(address))
	}
	// remote call
	err = (*cmd.DefaultOptions().Client).Call(ctx, req, &response, opts...)
	if err != nil {
		ce := errors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = "go.micro.rpc"
			ce.Status = http.StatusText(500)
			ce.Detail = "error during request: " + ce.Detail
			c.JSON(500, ce)
		default:
			c.JSON(int(ce.Code), ce)
		}
		return
	}
	// b, _ := response.MarshalJSON()
	c.JSON(200, response)
	// w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	// w.Write(b)

	// raw, err := c.GetRawData()
	// if err != nil {
	// 	c.String(500, "%s", err.Error())
	// 	return
	// }
	// if len(raw) == 0 {
	// 	c.String(500, "%s", "NO DATA")
	// 	return
	// }
	// // fmt.Println(raw)
	// paySrv := pay_srv.NewPayService("go.micro.srv.pay_srv", client.DefaultClient)
	// rst, err := paySrv.Create(c, &pay_srv.CreateReq{})
	// if err != nil {
	// 	c.String(500, "%s", err.Error())
	// 	return
	// }
	// // fmt.Println(rst.Result)
	// c.String(200, "%s", rst.Result)
}

//RequestToContext ..
func RequestToContext(r *http.Request) context.Context {
	ctx := context.Background()
	md := make(metadata.Metadata)
	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}
