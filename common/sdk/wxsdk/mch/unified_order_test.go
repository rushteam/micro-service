package mch

import (
	"testing"
)

func Test_Call(t *testing.T) {
	uo := &UnifiedOrderReq{}
	orderRsp, err := uo.Call()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(orderRsp)
	// t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	// t.Log("第一个测试通过了")    //记录一些你期望记录的信息
}
