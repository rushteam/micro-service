package repository

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/rushteam/gosql"
	"github.com/rushteam/micro-service/service/user-srv/model"
	// "upper.io/db.v3"
)

var (
	//ErrPassword 密码不正确
	ErrPassword = errors.New(`密码错误`)
	//ErrHasCreated 已注册
	ErrHasCreated = errors.New(`用户已注册`)
	//ErrUnvaildLoginType 无效的登陆方式
	ErrUnvaildLoginType = errors.New("无效的登陆方式")
	//loginMaps 本地登陆方式,即需要加密密码
	loginMaps = map[string]bool{"phone": true, "email": true, "username": true}
)

//userRepository ..
type userRepository struct{}

func newUserRepo() *userRepository {
	return &userRepository{}
}

//User ..
var User = newUserRepo()

//Create user
func (repo userRepository) CreateByPhone(user *model.UserModel, phone, pwd string) error {
	tx, err := model.DB.Begin()
	if err != nil {
		return err
	}
	now := time.Now()
	//user
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err = tx.Insert(user)

	if err != nil {
		tx.Rollback()
		return err
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	//login
	login := &model.LoginModel{}
	login.UID = user.UID
	login.Platform = "phone"
	login.Openid = phone
	login.AccessToken = pwd
	login.Status = 1
	login.CreatedAt = now
	login.UpdatedAt = now
	_, err = tx.Insert(login)
	if err != nil {
		tx.Rollback()
		if e, ok := err.(*gosql.Error); ok {
			if e.Number == 1062 {
				return ErrHasCreated
			}
		}
		return err
	}
	return tx.Commit()
}

//SigninByPwd ...
func (repo userRepository) SigninByPwd(platform, openid, password string) (*model.LoginModel, error) {
	login := &model.LoginModel{}
	err := model.DB.Fetch(
		login,
		gosql.Where("platform", platform),
		gosql.Where("openid", openid),
	)
	if login.AccessToken != password {
		return nil, ErrPassword
	}
	return login, err
}

//获取哈希密码
func getHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	str := hex.EncodeToString(hash.Sum(nil))
	return str
}

//LoginByPassword ...
// func (sess *LoginModel) LoginByPassword(typ, openid, pwd string) (*LoginModel, error) {
// 	var login LoginModel
// 	// return &login, nil
// 	result := sess.Where("platform = ?", typ).Where("openid = ?", openid).First(&login)
// 	if result.Error != nil {
// 		return nil, errors.New("用户不存在")
// 	}
// 	if login.AccessToken != pwd {
// 		return nil, errors.New("密码错误")
// 	}
// 	return &login, nil
// }

// //UserByUID ...
// func (sess *Session) UserByUID(uid int64) (*UserModel, error) {
// 	var user UserModel
// 	result := sess.Where("uid = ?", uid).First(&user)
// 	if result.Error != nil {
// 		return nil, errors.New("用户不存在")
// 	}
// 	return &user, nil
// }

// //UserAdd ..
// func (sess *Session) UserAdd(u *UserModel) error {
// 	result := sess.Create(u)
// 	if result.Error != nil {
// 		return errors.New("账户创建失败")
// 	}
// 	return nil
// }
