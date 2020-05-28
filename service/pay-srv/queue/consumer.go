package queue

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/micro/go-micro/v2/logger"
	"github.com/rushteam/micro-service/common/pb/pay_srv"
)

//Consumer 消费者
type Consumer struct{}

//Process  Method can be of any name
func (s *Consumer) Process(ctx context.Context, event *pay_srv.NotifyEvent) error {
	// md, _ := metadata.FromContext(ctx)
	logger.Infof("[Consumer.Process] recvied data: %+v\r\n", event)
	if event.GetUrl() == "" || event.GetBody() == "" || event.GetPayNo() == "" {
		logger.Infof("[Consumer.Process] notifyEvent.Data not empty")
		return fmt.Errorf("[Consumer.Process] notifyEvent.Data not empty")
	}
	paramsReader := bytes.NewBufferString(event.Body)
	resp, err := http.Post(event.Url, "application/json", paramsReader)
	if err != nil {
		logger.Errorf("[Consumer.Process] failed,post error: %s", err.Error())
		return fmt.Errorf("[Consumer.Process] notifyEvent.Data not empty")
	}
	defer resp.Body.Close()
	var state = false
	for {
		if resp.StatusCode != http.StatusOK {
			logger.Errorf("[Consumer.Process] failed,response status code: %d", resp.StatusCode)
			break
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("[Consumer.Process] failed,read %s", err)
			break
		}
		data := string(body)
		if data != "OK" {
			logger.Errorf("[Consumer.Process] failed,return %s", data)
			break
		}
		// state = true
		break
	}
	//更新状态
	_ = state
	/*
		t := &model.TradeModel{}
		if state == false {
			_, err = gosql.Model(t).UpdateField("[+]notify_num", 1).Where("pay_no", event.GetPayNo()).Update()
			if err != nil {
				logger.Logf("[Consumer.Process] ERROR update error %s", err.Error())
			}
			return fmt.Errorf("failed")
		}
		_, err = orm.Model(t).UpdateField("[+]notify_num", 1).UpdateField("[+]notify_state", 1).Where("pay_no", event.GetPayNo()).Update()
		if err != nil {
			logger.Logf("[Consumer.Process] INFO update error %s", err.Error())
		}
	*/
	return nil
}
