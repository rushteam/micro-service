package handler

import (
	"context"

	"gitee.com/rushteam/micro-service/common/pb/user_srv"
	// "go.uber.org/zap"
)

//UserServiceHandler ...
type UserServiceHandler struct {
	// logger *zap.Logger
}

//NewUserServiceHandler ...
func NewUserServiceHandler(ctx context.Context) *UserServiceHandler {
	return &UserServiceHandler{}
}

//Login ...
func (wx *UserServiceHandler) Login(ctx context.Context, req *user_srv.LoginReq, rsp *user_srv.LoginRsq) error {

	return nil
}
