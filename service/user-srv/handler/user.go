package handler

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/mlboy/micro-service/common/pb/usersrv"
	"github.com/mlboy/micro-service/common/utils"
	"github.com/mlboy/micro-service/service/user-srv/repository"
	"github.com/mlboy/micro-service/service/user-srv/session"

	"upper.io/db.v3/lib/sqlbuilder"
	// "go.uber.org/zap"
)

//RegisterUserServiceHandler ..
func RegisterUserServiceHandler(service micro.Service, d sqlbuilder.Database) {
	// usersrv.RegisterUserServiceHandler(service.Server(), NewUserService(d))
	usersrv.RegisterUserServiceHandler(service.Server(), NewUserService(d))
}

//NewUserService ..
func NewUserService(d sqlbuilder.Database) *UserService {
	return &UserService{d}
}

//UserService ...
type UserService struct {
	db sqlbuilder.Database
	// logger *zap.Logger
}

func validatePhone(phone string) bool {
	var regular = "^(((13[0-9])|(14[579])|(15([0-3]|[5-9]))|(16[6])|(17[0135678])|(18[0-9])|(19[89]))\\d{8})$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

//Login ...
func (s *UserService) Login(ctx context.Context, req *usersrv.LoginReq, rsp *usersrv.AuthRsp) error {
	log.Log("[access] UserService.Login")
	//phone or email or username
	if req.GetType() == "" {
		return errors.BadRequest("UserService.Login", "type参数不能为空")
	}
	if _, ok := repository.LocalLoginList[req.GetType()]; ok {
	} //账号登陆 本地登陆账号
	//登录名+密码登陆
	if req.GetType() == "phone" {
		if !validatePhone(req.Login) {
			return errors.BadRequest("UserService.Login", "手机号格式错误")
		}
		//密码位数不在登陆时候验证，而是在设置时候验证
		// if len(req.Password) < 6 { //密码不得小于6位
		// 	return errors.BadRequest("UserService.Login", "密码不得小于6位")
		// }
		//test --md5--> 098f6bcd4621d373cade4e832627b4f6
		loginRepo := &repository.LoginRepository{Db: s.db}
		login, err := loginRepo.FindByPassword(req.Type, req.Login, req.Password)
		if err != nil {
			return errors.BadRequest("UserService.Login", "用户名或密码错误")
		}
		rsp.Uid = login.UID
		// gen token
		subject := strconv.FormatInt(login.UID, 10)
		token := session.New("user-srv", subject, "")
		jwt, err := session.Encode(token, "")
		if err != nil {
			return errors.BadRequest("UserService.Login", "登录异常,请请联系客服")
		}
		rsp.Token = jwt
	} else {
		return errors.BadRequest("UserService.Login", "未知登陆方式")
	}
	return nil
}

//Register ..
// func (s *UserService) Register(ctx context.Context, req *usersrv.RegisterData, rsp *usersrv.UserData) error {
// 	log.Log("[access] UserService.Create")
// 	if len(req.LoginList) < 1 {
// 		return errors.BadRequest("UserService.Create", "注册失败,账号信息不全")
// 	}
// 	if req.GetUser() == nil {
// 		return errors.BadRequest("UserService.Create", "注册失败,用户信息不全")
// 	}
// 	//注册新用户逻辑
// 	if req.User.GetNickname() == "" {
// 		return errors.BadRequest("UserService.Create", "注册失败,昵称不能为空")
// 	}
// 	if req.User.GetNickname() == "" {
// 		return errors.BadRequest("UserService.Create", "注册失败,昵称不能为空")
// 	}
// 	var userData repository.UserModel
// 	tx, err := s.db.NewTx(ctx)
// 	userRepo := &repository.UserRepository{Db: tx}
// 	userData.Nickname = req.User.Nickname
// 	// user. = req.User.Firstname
// 	// user.Lastname = req.User.Lastname
// 	userData.Gender = req.User.Gender
// 	userData.Avatar = req.User.Avatar
// 	user, err := userRepo.Create(userData)
// 	if err != nil {
// 		tx.Rollback()
// 		log.Log("创建新用户失败" + err.Error())
// 		return errors.BadRequest("UserService.Create", "注册失败,请重试")
// 	}
// 	loginRepo := &repository.LoginRepository{Db: tx}
// 	for _, login := range req.LoginList {
// 		if login.GetPlatform() == "" {
// 			return errors.BadRequest("UserService.Create", "注册失败,登陆类别不能为空")
// 		}
// 		if login.GetLogin() == "" {
// 			return errors.BadRequest("UserService.Create", "注册失败,账号ID不能为空")
// 		}
// 		if login.GetPassword() == "" {
// 			return errors.BadRequest("UserService.Create", "注册失败,账号凭证不能为空")
// 		}
// 		loginData := repository.LoginModel{}
// 		loginData.UID = user.UID
// 		loginData.Platform = login.Platform
// 		loginData.Openid = login.Login
// 		loginData.AccessToken = login.Password
// 		loginData.AccessExpire = time.Now().Add(time.Hour * 24 * 7)
// 		_, err = loginRepo.Create(loginData)
// 		if err != nil {
// 			log.Log("login数据创建失败" + err.Error())
// 			tx.Rollback()
// 			return errors.BadRequest("UserService.Create", "注册失败,账号已经存在")
// 		}
// 	}
// 	tx.Commit()
// 	return nil
// }

//User ..
func (s *UserService) User(ctx context.Context, req *usersrv.UserRsp, rsp *usersrv.UserRsp) error {
	log.Log("[access] UserService.User")
	// Model := model.Db()
	token, err := session.Decode(req.GetJwt(), "")
	if err != nil {
		return errors.BadRequest("UserService.Login", "登录超时或TOKEN非法")
	}
	if token.Subject == "" || token.Subject == "0" {
		return errors.BadRequest("UserService.Login", "当前TOKEN未绑定用户")
	}
	uid, err := strconv.ParseInt(token.Subject, 10, 64)
	if err != nil {
		return errors.BadRequest("UserService.Login", "当前TOKEN无法解析用户")
	}
	userRepo := &repository.UserRepository{Db: s.db}
	user, err := userRepo.FindByUID(uid)
	if err != nil {
		return errors.BadRequest("UserService.Login", "用户不存在或已被锁定")
	}
	rsp.User = &usersrv.UserData{}
	rsp.User.Uid = user.UID
	rsp.User.Nickname = user.Nickname
	rsp.User.Gender = user.Gender
	rsp.User.Avatar = user.Avatar
	rsp.User.CreatedAt = utils.FormatDate(user.CreatedAt)
	rsp.User.UpdatedAt = utils.FormatDate(user.UpdatedAt)
	return nil
}

//Bind ...
func (s *UserService) Bind(ctx context.Context, req *user_srv.CreateReq, rsp *user_srv.UserRsp) error {
	log.Log("[access] UserService.Bind")
	if req.GetToken() == nil {
		return errors.BadRequest("UserService.Create", "绑定失败,当前状态未登录")
	}
	if len(req.LoginList) < 1 {
		return errors.BadRequest("UserService.Create", "绑定失败,账号信息不全")
	}

	var userData repository.UserModel
	tx, err := s.db.NewTx(ctx)
	if req.User.GetUid() != 0 {
		//注册新用户逻辑
		if req.User.GetNickname() == "" {
			return errors.BadRequest("UserService.Create", "注册失败,昵称不能为空")
		}
		if req.User.GetNickname() == "" {
			return errors.BadRequest("UserService.Create", "注册失败,昵称不能为空")
		}
		userRepo := &repository.UserRepository{Db: tx}
		userData.Nickname = req.User.Nickname
		// user. = req.User.Firstname
		// user.Lastname = req.User.Lastname
		userData.Gender = req.User.Gender
		userData.Avatar = req.User.Avatar
		user, err := userRepo.Create(userData)
		if err != nil {
			tx.Rollback()
			log.Log("创建新用户失败" + err.Error())
			return errors.BadRequest("UserService.Create", "注册失败,请重试")
		}
		userData.UID = user.UID
	} else {
		//绑定用户逻辑
		userRepo := &repository.UserRepository{Db: tx}
		user, err := userRepo.FindByUID(req.User.GetUid())
		if err != nil {
			tx.Rollback()
			log.Log("未找到当前用户" + err.Error())
		}
		userData.UID = user.UID
	}
	loginRepo := &repository.LoginRepository{Db: tx}
	for _, login := range req.LoginList {
		if login.GetPlatform() == "" {
			return errors.BadRequest("UserService.Create", "注册失败,登陆类别不能为空")
		}
		if login.GetLogin() == "" {
			return errors.BadRequest("UserService.Create", "注册失败,账号ID不能为空")
		}
		if login.GetPassword() == "" {
			return errors.BadRequest("UserService.Create", "注册失败,账号凭证不能为空")
		}
		loginData := repository.LoginModel{}
		loginData.UID = userData.UID
		loginData.Platform = login.Platform
		loginData.Openid = login.Login
		loginData.AccessToken = login.Password
		loginData.AccessExpire = time.Now().Add(time.Hour * 24 * 7)
		_, err = loginRepo.Create(loginData)
		if err != nil {
			log.Log("login数据创建失败" + err.Error())
			tx.Rollback()
			return errors.BadRequest("UserService.Create", "注册失败,账号已经存在")
		}
	}
	tx.Commit()
	return nil
	return nil
}

//UnBind ...
func (s *UserService) UnBind(ctx context.Context, req *user_srv.UserReq, rsp *user_srv.UserRsp) error {
	return nil
}

//Update ...
func (s *UserService) Update(ctx context.Context, req *user_srv.UpdateReq, rsp *user_srv.UserRsp) error {
	return nil
}

//find -r "*.php" -exec 'cat' {} \; > /tmp/code.txt
