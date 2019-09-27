package handler

import (
	"context"
	"regexp"
	"strconv"

	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/mlboy/micro-service/common/pb/usersrv"
	"github.com/mlboy/micro-service/common/sdk/wxsdk"
	"github.com/mlboy/micro-service/common/utils"
	"github.com/mlboy/micro-service/service/user-srv/repository"
	"github.com/mlboy/micro-service/service/user-srv/session"

	"upper.io/db.v3/lib/sqlbuilder"
	// "go.uber.org/zap"
)

//RegisterUserServiceHandler ..
func RegisterUserServiceHandler(service micro.Service, d sqlbuilder.Database) {
	usersrv.RegisterUserServiceHandler(service.Server(), &UserService{db: d})
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

//LoginByPassword ...
func (s *UserService) LoginByPassword(ctx context.Context, req *usersrv.LoginByPasswordReq, rsp *usersrv.AuthRsp) error {
	log.Log("[access] UserService.LoginByPassword")
	//phone or email or username
	// if req.GetPlatform() == "" {
	// 	return errors.BadRequest("UserService.Login", "plattform参数不能为空")
	// }
	// // if _, ok := repository.LocalLoginList[req.GetPlatform()]; ok {
	// // } //账号登陆 本地登陆账号
	// //登录名+密码登陆
	// if req.GetPlatform() == "phone" {
	// 		} else {
	// 	return errors.BadRequest("UserService.Login", "未知登陆方式")
	// }
	//密码位数不在登陆时候验证，而是在设置时候验证
	// if len(req.Password) < 6 { //密码不得小于6位
	// 	return errors.BadRequest("UserService.Login", "密码不得小于6位")
	// }
	if !validatePhone(req.GetLoginname()) {
		return errors.BadRequest("UserService.LoginByPassword", "手机号格式错误")
	}
	loginRepo := &repository.LoginRepository{Db: s.db}
	login, err := loginRepo.FindByPassword("phone", req.GetLoginname(), req.GetPassword())
	if err != nil {
		return errors.BadRequest("UserService.LoginByPassword", "用户名或密码错误")
	}
	rsp.Uid = login.UID
	// gen token
	jwt, err := GenToken(login.UID)
	if err != nil {
		return errors.BadRequest("UserService.LoginByPassword", "登录异常,请请联系客服")
	}
	rsp.Token = jwt
	return nil
}

//LoginByOAuth ...
func (s *UserService) LoginByOAuth(ctx context.Context, req *usersrv.LoginByOAuthReq, rsp *usersrv.AuthRsp) error {
	log.Log("[access] UserService.LoginByOAuth")
	//phone or email or username
	if req.GetPlatform() == "" {
		return errors.BadRequest("UserService.LoginByOAuth", "platform参数不能为空")
	}
	if req.GetAppid() == "" {
		return errors.BadRequest("UserService.LoginByOAuth", "appid参数不能为空")
	}
	if req.GetSercet() == "" {
		return errors.BadRequest("UserService.LoginByOAuth", "sercet参数不能为空")
	}
	if req.GetCode() == "" {
		return errors.BadRequest("UserService.LoginByOAuth", "code参数不能为空")
	}
	// if req.GetOpenid() == "" {
	// 	return errors.BadRequest("UserService.LoginByOAuth", "openid参数不能为空")
	// }
	// if req.GetAccessToken() == "" {
	// 	return errors.BadRequest("UserService.LoginByOAuth", "access_token参数不能为空")
	// }
	if req.GetPlatform() == "wx" {
		//通过code 获取信息
		//accesstoken
		at, err := wxsdk.GetAuthAccessToken(ctx, req.GetAppid(), req.GetSercet(), req.GetCode())
		if err != nil {
			log.Logf(err.Error())
			return errors.BadRequest("UserService.LoginByOAuth", "请求第三方失败获取access_token失败")
		}
		//userinfo
		ui, err := wxsdk.GetUserinfo(ctx, at.AccessToken, at.OpenID)
		loginRepo := &repository.LoginRepository{Db: s.db}
		login, err := loginRepo.FindByPassword("wx_open_id", ui.OpenID, at.AccessToken)
		// login, err := loginRepo.FindByPassword("wx_union_id", ui.OpenID, at.AccessToken)
		if err != nil {
			return errors.BadRequest("UserService.Login", "用户名或密码错误")
		}
		rsp.Uid = login.UID
		// gen token
		jwt, err := GenToken(login.UID)
		if err != nil {
			return errors.BadRequest("UserService.Login", "登录异常,请请联系客服")
		}
		rsp.Token = jwt
	}
	return nil
}

//LoginByCaptcha ...
func (s *UserService) LoginByCaptcha(ctx context.Context, req *usersrv.LoginByCaptchaReq, rsp *usersrv.AuthRsp) error {
	log.Log("[access] UserService.LoginByCaptcha")
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
func (s *UserService) User(ctx context.Context, req *usersrv.UserReq, rsp *usersrv.UserRsp) error {
	log.Log("[access] UserService.User")
	// Model := model.Db()
	token, err := session.Decode(req.GetToken(), "")
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
	rsp.Uid = user.UID
	rsp.Nickname = user.Nickname
	rsp.Gender = user.Gender
	rsp.Avatar = user.Avatar
	rsp.CreatedAt = utils.FormatDate(user.CreatedAt)
	rsp.UpdatedAt = utils.FormatDate(user.UpdatedAt)
	return nil
}

//Bind ...
func (s *UserService) Bind(ctx context.Context, req *usersrv.BindReq, rsp *usersrv.UserRsp) error {
	log.Log("[access] UserService.Bind")
	if req.GetToken() == "" {
		return errors.BadRequest("UserService.Create", "绑定失败,当前状态未登录")
	}
	// if req.GetLogin() == nil {
	// 	return errors.BadRequest("UserService.Create", "绑定失败,缺少登陆信息")
	// }
	// token, err := session.Decode(req.GetToken(), "")
	// if err != nil {
	// 	return errors.BadRequest("UserService.Login", "登录超时或TOKEN非法")
	// }
	// if token.Subject == "" || token.Subject == "0" {
	// 	return errors.BadRequest("UserService.Login", "当前TOKEN未绑定用户")
	// }
	// uid, err := strconv.ParseInt(token.Subject, 10, 64)
	// if err != nil {
	// 	return errors.BadRequest("UserService.Login", "当前TOKEN无法解析用户")
	// }
	// tx, err := s.db.NewTx(ctx)
	// //绑定用户逻辑
	// userRepo := &repository.UserRepository{Db: tx}
	// user, err := userRepo.FindByUID(uid)
	// if err != nil {
	// 	log.Log("未找到当前用户" + err.Error())
	// }
	// loginRepo := &repository.LoginRepository{Db: tx}
	// if req.GetLogin().GetPlatform() == "" {
	// 	return errors.BadRequest("UserService.Create", "注册失败,登陆类别不能为空")
	// }
	// if req.GetLogin().GetLoginname() == "" {
	// 	return errors.BadRequest("UserService.Create", "注册失败,账号ID不能为空")
	// }
	// if req.GetLogin().GetPassword() == "" {
	// 	return errors.BadRequest("UserService.Create", "注册失败,账号凭证不能为空")
	// }
	// loginData := repository.LoginModel{}
	// loginData.UID = user.UID
	// loginData.Platform = req.GetLogin().GetPlatform()
	// loginData.Openid = req.GetLogin().GetLoginname()
	// loginData.AccessToken = req.GetLogin().GetPassword()
	// loginData.AccessExpire = time.Now().Add(time.Hour * 24 * 7 * 2)
	// _, err = loginRepo.Create(loginData)
	// if err != nil {
	// 	log.Log("login数据创建失败" + err.Error())
	// 	tx.Rollback()
	// 	return errors.BadRequest("UserService.Create", "注册失败,账号已经存在")
	// }
	// tx.Commit()
	return nil
}

//Unbind ...
func (s *UserService) Unbind(ctx context.Context, req *usersrv.UnbindReq, rsp *usersrv.UserRsp) error {
	return nil
}

//Update ...
func (s *UserService) Update(ctx context.Context, req *usersrv.UpdateReq, rsp *usersrv.UserRsp) error {
	return nil
}

//GenToken 生成token
func GenToken(uid int64) (string, error) {
	subject := strconv.FormatInt(uid, 10)
	token := session.New("user-srv", subject, "")
	jwt, err := session.Encode(token, "")
	return jwt, err
}

//find -r "*.php" -exec 'cat' {} \; > /tmp/code.txt
