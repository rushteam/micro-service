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
	if event.Url == "" || event.Body == "" {
		log.Logf("[Consumer.Process] notifyEvent.Data not empty")
	}
	// statusCode, body, err := utils.HttpPost(url, []byte(event.Message))
	paramsReader := bytes.NewBufferString(event.Body)
	resp, err := http.Post(event.Url, "application/json", paramsReader)
	if err != nil {
		log.Logf("[Consumer.Process] failed,post error: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Logf("[Consumer.Process] failed,response status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	data := string(body)
	fmt.Println(string(body))
	if data != "OK" {
		return fmt.Errorf("[Consumer.Process] failed,return %s", data)
	}
	// todo 更新状态
	t := model.TradeModel{}
	// orm.Model(t).Update(func(s *builder.SQLSegments) {
	// 	s.Where("")
	// })
	// orm.Model(t).UpdateField("[+]notify_num", 1).Update()
	// orm.Exec("")
	// orm.Model(t).Update()
	_, err = orm.Model(t).UpdateField("[+]notify_num", 1).Where("pvd_out_trade_no", event.Id).Update()
	if err != nil {
		return err
	}
	return nil
}
