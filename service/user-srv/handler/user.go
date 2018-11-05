package handler

import (
	"context"
	"regexp"

	"github.com/micro/go-micro/errors"

	"gitee.com/rushteam/micro-service/common/utils"

	"gitee.com/rushteam/micro-service/service/user-srv/model"

	"gitee.com/rushteam/micro-service/common/pb/user_srv"
	"github.com/micro/go-log"
	// "go.uber.org/zap"
)

//UserService ...
type UserService struct {
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

var localLoginList = []string{"phone", "email", "username"}

//Login ...
func (s *UserService) Login(ctx context.Context, req *user_srv.LoginReq, rsp *user_srv.LoginRsq) error {
	log.Log("[access] UserService.Login")
	//phone or email or username
	Model := model.Db()
	if utils.SliceIndexOf(req.Platform, localLoginList) >= 0 { //账号登陆
		if req.Platform == "phone" && !validatePhone(req.Openid) {
			// return errors.New("手机号格式错误")
			return errors.BadRequest("UserService.Login", "手机号格式错误")
		}
		if len(req.AccessToken) < 6 { //密码不得小于6位
			// return errors.New("密码错误")
			return errors.BadRequest("UserService.Login", "密码错误")
		}
		login, err := Model.LoginByPassword(req.Platform, req.Openid, req.AccessToken)
		if err != nil {
			return errors.BadRequest("UserService.Login", "用户名或密码错误")
			// return errors.New("用户名或密码错误")
		}
		// fmt.Println(login)
		rsp.Uid = login.UID
	} else { //三方登陆

	}
	return nil
}

//User ..
func (s *UserService) User(ctx context.Context, req *user_srv.UserReq, rsp *user_srv.UserRsq) error {
	log.Log("[access] UserService.User")
	Model := model.Db()
	user, err := Model.UserByUID(req.Uid)
	if err != nil {
		return errors.BadRequest("UserService.Login", "用户名不存在")
		// return errors.New("用户名不存在")
	}
	rsp.Userinfo.Uid = user.UID
	rsp.Userinfo.Nickname = user.Nickname
	rsp.Userinfo.Gender = user.Gender
	rsp.Userinfo.Avatar = user.Avatar
	rsp.Userinfo.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	rsp.Userinfo.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")
	return nil
}

//Create ..
func (s *UserService) Create(ctx context.Context, req *user_srv.CreateReq, rsp *user_srv.UserRsq) error {
	log.Log("[access] UserService.Create")
	if len(req.LoginList) < 1 {
		return errors.BadRequest("UserService.Create", "注册失败,账号信息不全")
	}
	if req.GetUserinfo() == nil {
		return errors.BadRequest("UserService.Create", "注册失败,用户信息不全")
	}
	if req.Userinfo.GetNickname() == "" {
		return errors.BadRequest("UserService.Create", "注册失败,昵称不能为空")
	}
	Model := model.Begin()
	// req.Userinfo
	var u = &model.UserModel{}
	u.Nickname = req.GetUserinfo().Nickname
	u.Avatar = req.Userinfo.Avatar
	u.Gender = req.Userinfo.Gender
	err := Model.UserAdd(u)
	if err != nil {
		return errors.BadRequest("UserService.Create", "注册失败,请重试")
	}
	for _, login := range req.LoginList {
		if utils.SliceIndexOf(login.Platform, localLoginList) >= 0 { //账号登陆
			_, err = Model.LoginAdd(
				req.Userinfo.Uid,
				login.Platform,
				login.Openid,
				login.AccessToken)
			if err !=nil {
				Model.Callback()
				return errors.BadRequest("UserService.Create", "注册失败,账号已经存在"+err.Error())
			}
		} else {  //三方登录
		}
	}

	return nil
}
