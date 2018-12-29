package queue

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	log "github.com/micro/go-log"

	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	// "github.com/go-log/log"
)

//Consumer 消费者
type Consumer struct{}

//Process  Method can be of any name
func (s *Consumer) Process(ctx context.Context, event *pay_srv.NotifyEvent) error {
	// md, _ := metadata.FromContext(ctx)
	log.Logf("[Consumer.Process] recvied data: %+v\r\n", event)
	if event.Data == nil || event.Data.Url == "" || event.Data.Body == "" {
		log.Logf("[Consumer.Process] notifyEvent.Data not empty")
	}
	var url = "http://1thx.com/"
	// statusCode, body, err := utils.HttpPost(url, []byte(event.Message))
	paramsReader := bytes.NewBufferString(event.Data.Body)
	resp, err := http.Post(event.Data.Url, "application/json", paramsReader)
	if err != nil {
		log.Logf("[Consumer.Process] post error: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Logf("[Consumer.Process] response status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if body != "OK" {
		return log.Logf("[Consumer.Process] post error %s", err.Error())
	} else {

	}
}
