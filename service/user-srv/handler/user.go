package handler

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"gitee.com/rushteam/micro-service/service/user-srv/model"

	"gitee.com/rushteam/micro-service/common/pb/user_srv"
	"github.com/micro/go-log"
	// "go.uber.org/zap"
)

//UserServiceHandler ...
type UserServiceHandler struct {
	// logger *zap.Logger
}

//NewUserServiceHandler ...
// func NewUserServiceHandler(ctx context.Context) *UserServiceHandler {
// 	return &UserServiceHandler{}
// }
func validatePhone(phone string) bool {
	var regular = "^(((13[0-9])|(14[579])|(15([0-3]|[5-9]))|(16[6])|(17[0135678])|(18[0-9])|(19[89]))\\d{8})$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

//Login ...
func (wx *UserServiceHandler) Login(ctx context.Context, req *user_srv.LoginReq, rsp *user_srv.LoginRsq) error {
	log.Log("[access] UserServiceHandler.Login")
	//phone or email or username
	if req.Platform == "phone" { //手机登陆
		if !validatePhone(req.Openid) {
			return errors.New("手机号格式错误")
		}
		if len(req.AccessToken) < 6 { //密码不得小于6位
			return errors.New("密码错误")
		}
		login := model.LoginByPhone(req.Openid, req.AccessToken)
		fmt.Println(login)
	} else { //三方登陆

	}
	return nil
}
