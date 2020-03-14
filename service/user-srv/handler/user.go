package handler

import (
	"context"
	"regexp"
	"strconv"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/rushteam/micro-service/common/pb/usersrv"
	"github.com/rushteam/micro-service/service/user-srv/model"
	"github.com/rushteam/micro-service/service/user-srv/repository"
	"github.com/rushteam/micro-service/service/user-srv/session"
	// "go.uber.org/zap"
)

//RegisterUserServiceHandler ..
func RegisterUserServiceHandler(srv micro.Service) {
	usersrv.RegisterUserServiceHandler(srv.Server(), &UserService{
		auth: srv.Options().Auth,
	})
}

//UserService ...
type UserService struct {
	auth auth.Auth
	// logger *zap.Logger
}

//验证手机号
func validatePhone(phone string) bool {
	var regular = "^(((13[0-9])|(14[579])|(15([0-3]|[5-9]))|(16[6])|(17[0135678])|(18[0-9])|(19[89]))\\d{8})$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

//GenToken 生成token
func GenToken(uid int64) (string, error) {
	subject := strconv.FormatInt(uid, 10)
	token := session.New("user-srv", subject, "")
	jwt, err := session.Encode(token, "")
	return jwt, err
}

//Signin 登陆(手机号+密码)
func (s *UserService) Signin(ctx context.Context, req *usersrv.SigninReq, rsp *usersrv.AuthRsp) error {
	log.Tracef("[access] UserService.Signin")
	//密码位数不在登陆时候验证，而是在设置时候验证
	if !validatePhone(req.GetLoginname()) {
		return errors.BadRequest("UserService.Signin", "手机号格式错误")
	}
	login, err := repository.User.SigninByPwd("phone", req.GetLoginname(), req.GetPassword())
	if err != nil {
		return errors.BadRequest("UserService.Signin", "账号或密码错误")
	}
	rsp.Uid = login.UID
	// Generate an auth account
	acc, err := s.auth.Generate(strconv.FormatInt(login.UID, 10))
	if err != nil {
		return errors.InternalServerError("UserService.Signin", "登录异常,请联系客服(%v)", err)
	}
	rsp.Token = acc.Token

	// gen token
	// jwt, err := GenToken(login.UID)
	// if err != nil {
	// 	return errors.BadRequest("UserService.Signin", "登录异常,请联系客服")
	// }
	// rsp.Token = jwt
	return nil
}

//SigninByPhoneCaptcha 手机号+验证码
func (s *UserService) SigninByPhoneCaptcha(ctx context.Context, req *usersrv.SigninByPhoneCaptchaReq, rsp *usersrv.AuthRsp) error {
	log.Tracef("[access] UserService.SigninByPhoneCaptcha")
	return nil
}

//Signup 注册用户
func (s *UserService) Signup(ctx context.Context, req *usersrv.SignupReq, rsp *usersrv.UserInfo) error {
	log.Tracef("[access] UserService.Signup")
	//注册新用户逻辑
	if req.GetNickname() == "" {
		return errors.BadRequest("UserService.Create", "注册失败,昵称不能为空")
	}
	if req.GetPhone() == "" {
		return errors.BadRequest("UserService.Create", "注册失败,手机号不能为空")
	}
	//todo 验证手机号 和 验证码 是否正确 (phone + vcode)
	user := &model.UserModel{}
	user.Nickname = req.GetNickname()
	user.Gender = req.GetGender()
	user.Avatar = req.GetAvatar()
	user.Status = 1
	err := repository.User.CreateByPhone(user, req.GetPhone(), req.GetPassword())
	if err != nil {
		return errors.BadRequest("UserService.Create", "注册失败,%s", err.Error())
	}
	rsp.Uid = user.UID
	rsp.Nickname = req.GetNickname()
	rsp.Gender = req.GetGender()
	rsp.Avatar = req.GetAvatar()
	rsp.Phone = req.GetPhone()
	rsp.Status = user.Status
	rsp.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	rsp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")

	// string phone = 8;//手机号
	// string email = 9;//邮箱
	return nil
}

//OAuthAuthorize ..
func (s *UserService) OAuthAuthorize(ctx context.Context, req *usersrv.OAuthAuthorizeReq, rsp *usersrv.OAuthAuthorizeRsp) error {
	log.Tracef("[access] UserService.OAuthAuthorize")
	return nil
}

//Register ..
// func (s *UserService) Register(ctx context.Context, req *usersrv.RegisterData, rsp *usersrv.UserData) error {
// 	log.Log("[access] UserService.Create")
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

//User 获取用户信息
// func (s *UserService) User(ctx context.Context, req *usersrv.UserReq, rsp *usersrv.UserRsp) error {
// 	log.Log("[access] UserService.User")
// 	// Model := model.Db()
// 	token, err := session.Decode(req.GetToken(), "")
// 	if err != nil {
// 		return errors.BadRequest("UserService.Login", "登录超时或TOKEN非法")
// 	}
// 	if token.Subject == "" || token.Subject == "0" {
// 		return errors.BadRequest("UserService.Login", "当前TOKEN未绑定用户")
// 	}
// 	uid, err := strconv.ParseInt(token.Subject, 10, 64)
// 	if err != nil {
// 		return errors.BadRequest("UserService.Login", "当前TOKEN无法解析用户")
// 	}
// 	userRepo := &repository.UserRepository{Db: s.db}
// 	user, err := userRepo.FindByUID(uid)
// 	if err != nil {
// 		return errors.BadRequest("UserService.Login", "用户不存在或已被锁定")
// 	}
// 	rsp.Uid = user.UID
// 	rsp.Nickname = user.Nickname
// 	rsp.Gender = user.Gender
// 	rsp.Avatar = user.Avatar
// 	rsp.CreatedAt = utils.FormatDate(user.CreatedAt)
// 	rsp.UpdatedAt = utils.FormatDate(user.UpdatedAt)
// 	return nil
// }

// //Update 更新字段
// func (s *UserService) Update(ctx context.Context, req *usersrv.UpdateReq, rsp *usersrv.UserRsp) error {
// 	return nil
// }

// //LoginByOAuth oauth2 code登陆
// func (s *UserService) LoginByOAuth(ctx context.Context, req *usersrv.LoginByOAuthReq, rsp *usersrv.AuthRsp) error {
// 	log.Log("[access] UserService.LoginByOAuth")
// 	//phone or email or username
// 	if req.GetPlatform() == "" {
// 		return errors.BadRequest("UserService.LoginByOAuth", "platform参数不能为空")
// 	}
// 	if req.GetAppid() == "" {
// 		return errors.BadRequest("UserService.LoginByOAuth", "appid参数不能为空")
// 	}
// 	if req.GetSercet() == "" {
// 		return errors.BadRequest("UserService.LoginByOAuth", "sercet参数不能为空")
// 	}
// 	if req.GetCode() == "" {
// 		return errors.BadRequest("UserService.LoginByOAuth", "code参数不能为空")
// 	}
// 	// if req.GetOpenid() == "" {
// 	// 	return errors.BadRequest("UserService.LoginByOAuth", "openid参数不能为空")
// 	// }
// 	// if req.GetAccessToken() == "" {
// 	// 	return errors.BadRequest("UserService.LoginByOAuth", "access_token参数不能为空")
// 	// }
// 	if req.GetPlatform() == "wx" {
// 		//通过code 获取信息
// 		//accesstoken
// 		at, err := wxsdk.GetAuthAccessToken(ctx, req.GetAppid(), req.GetSercet(), req.GetCode())
// 		if err != nil {
// 			log.Logf(err.Error())
// 			return errors.BadRequest("UserService.LoginByOAuth", "请求第三方失败获取access_token失败")
// 		}
// 		//userinfo
// 		ui, err := wxsdk.GetUserinfo(ctx, at.AccessToken, at.OpenID)
// 		loginRepo := &repository.LoginRepository{Db: s.db}
// 		var login *repository.LoginModel
// 		if ui.Unionid == "" {
// 			login, err = loginRepo.FindByPassword("wx_open_id", ui.OpenID, at.AccessToken)
// 		} else {
// 			login, err = loginRepo.FindByPassword("wx_union_id", ui.Unionid, at.AccessToken)
// 		}
// 		if err != nil {
// 			log.Logf(err.Error())
// 			return errors.BadRequest("UserService.LoginByOAuth", "当前用户未注册")
// 		}
// 		//自动注册逻辑
// 		rsp.Uid = login.UID
// 		// gen token
// 		jwt, err := GenToken(login.UID)
// 		if err != nil {
// 			return errors.BadRequest("UserService.LoginByOAuth", "登录异常,请请联系客服")
// 		}
// 		rsp.Token = jwt
// 	}
// 	if req.GetPlatform() == "wxa" {
// 		//通过code 获取信息
// 		at, err := wxsdk.GetCode2Session(ctx, req.GetAppid(), req.GetSercet(), req.GetCode())
// 		if err != nil {
// 			log.Logf(err.Error())
// 			return errors.BadRequest("UserService.LoginByOAuth", "请求第三方失败获取access_token失败")
// 		}
// 		//userinfo
// 		ui, err := wxsdk.GetUserinfo(ctx, at.AccessToken, at.OpenID)
// 		loginRepo := &repository.LoginRepository{Db: s.db}
// 		var login *repository.LoginModel
// 		if ui.Unionid == "" {
// 			login, err = loginRepo.FindByPassword("wx_open_id", ui.OpenID, at.AccessToken)
// 		} else {
// 			login, err = loginRepo.FindByPassword("wx_union_id", ui.Unionid, at.AccessToken)
// 		}
// 		if err != nil {
// 			log.Logf(err.Error())
// 			return errors.BadRequest("UserService.LoginByOAuth", "当前用户未注册")
// 		}
// 		rsp.Uid = login.UID
// 		// gen token
// 		jwt, err := GenToken(login.UID)
// 		if err != nil {
// 			return errors.BadRequest("UserService.LoginByOAuth", "登录异常,请请联系客服")
// 		}
// 		rsp.Token = jwt
// 	}
// 	return nil
// }

// //Bind ...
// func (s *UserService) Bind(ctx context.Context, req *usersrv.BindReq, rsp *usersrv.UserRsp) error {
// 	log.Log("[access] UserService.Bind")
// 	if req.GetToken() == "" {
// 		return errors.BadRequest("UserService.Create", "绑定失败,当前状态未登录")
// 	}
// 	// if req.GetLogin() == nil {
// 	// 	return errors.BadRequest("UserService.Create", "绑定失败,缺少登陆信息")
// 	// }
// 	// token, err := session.Decode(req.GetToken(), "")
// 	// if err != nil {
// 	// 	return errors.BadRequest("UserService.Login", "登录超时或TOKEN非法")
// 	// }
// 	// if token.Subject == "" || token.Subject == "0" {
// 	// 	return errors.BadRequest("UserService.Login", "当前TOKEN未绑定用户")
// 	// }
// 	// uid, err := strconv.ParseInt(token.Subject, 10, 64)
// 	// if err != nil {
// 	// 	return errors.BadRequest("UserService.Login", "当前TOKEN无法解析用户")
// 	// }
// 	// tx, err := s.db.NewTx(ctx)
// 	// //绑定用户逻辑
// 	// userRepo := &repository.UserRepository{Db: tx}
// 	// user, err := userRepo.FindByUID(uid)
// 	// if err != nil {
// 	// 	log.Log("未找到当前用户" + err.Error())
// 	// }
// 	// loginRepo := &repository.LoginRepository{Db: tx}
// 	// if req.GetLogin().GetPlatform() == "" {
// 	// 	return errors.BadRequest("UserService.Create", "注册失败,登陆类别不能为空")
// 	// }
// 	// if req.GetLogin().GetLoginname() == "" {
// 	// 	return errors.BadRequest("UserService.Create", "注册失败,账号ID不能为空")
// 	// }
// 	// if req.GetLogin().GetPassword() == "" {
// 	// 	return errors.BadRequest("UserService.Create", "注册失败,账号凭证不能为空")
// 	// }
// 	// loginData := repository.LoginModel{}
// 	// loginData.UID = user.UID
// 	// loginData.Platform = req.GetLogin().GetPlatform()
// 	// loginData.Openid = req.GetLogin().GetLoginname()
// 	// loginData.AccessToken = req.GetLogin().GetPassword()
// 	// loginData.AccessExpire = time.Now().Add(time.Hour * 24 * 7 * 2)
// 	// _, err = loginRepo.Create(loginData)
// 	// if err != nil {
// 	// 	log.Log("login数据创建失败" + err.Error())
// 	// 	tx.Rollback()
// 	// 	return errors.BadRequest("UserService.Create", "注册失败,账号已经存在")
// 	// }
// 	// tx.Commit()
// 	return nil
// }

// //Unbind 解绑手机号
// func (s *UserService) Unbind(ctx context.Context, req *usersrv.UnbindReq, rsp *usersrv.UserRsp) error {
// 	return nil
// }
