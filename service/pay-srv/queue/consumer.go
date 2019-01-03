package queue

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitee.com/rushteam/micro-service/service/pay-srv/model"

	log "github.com/micro/go-log"
	"github.com/mlboy/godb/orm"

	"gitee.com/rushteam/micro-service/common/pb/pay_srv"
	// "github.com/go-log/log"
)

//Consumer 消费者
type Consumer struct{}

//Process  Method can be of any name
func (s *Consumer) Process(ctx context.Context, event *pay_srv.NotifyEvent) error {
	// md, _ := metadata.FromContext(ctx)
	log.Logf("[Consumer.Process] recvied data: %+v\r\n", event)
	if event.GetUrl() == "" || event.GetBody() == "" || event.GetPayNo() == "" {
		log.Logf("[Consumer.Process] notifyEvent.Data not empty")
		return fmt.Errorf("[Consumer.Process] notifyEvent.Data not empty")
	}
	paramsReader := bytes.NewBufferString(event.Body)
	resp, err := http.Post(event.Url, "application/json", paramsReader)
	if err != nil {
		log.Logf("[Consumer.Process] failed,post error: %s", err.Error())
		return fmt.Errorf("[Consumer.Process] notifyEvent.Data not empty")
	}
	defer resp.Body.Close()
	var state = false
	for {
		if resp.StatusCode != http.StatusOK {
			log.Logf("[Consumer.Process] failed,response status code: %d", resp.StatusCode)
			break
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Logf("[Consumer.Process] failed,read %s", err)
			break
		}
		data := string(body)
		if data != "OK" {
			log.Logf("[Consumer.Process] failed,return %s", data)
			break
		}
		state = true
		break
	}
	//更新状态
	t := &model.TradeModel{}
	if state == false {
		_, err = orm.Model(t).UpdateField("[+]notify_num", 1).Where("pay_no", event.GetPayNo()).Update()
		if err != nil {
			log.Logf("[Consumer.Process] ERROR update error %s", err.Error())
		}
		return fmt.Errorf("failed")
	}
	_, err = orm.Model(t).UpdateField("[+]notify_num", 1).UpdateField("[+]notify_state", 1).Where("pay_no", event.GetPayNo()).Update()
	if err != nil {
		log.Logf("[Consumer.Process] INFO update error %s", err.Error())
	}
	return nil
}

// orm.Model(t).Update(func(s *builder.SQLSegments) {
// 	s.Where("")
// })
// orm.Model(t).UpdateField("[+]notify_num", 1).Update()
// orm.Exec("")
// orm.Model(t).Update()
