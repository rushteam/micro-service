package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//LoginModel ..
type LoginModel struct {
	// gorm.Model
	ID           int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Platform     string
	Openid       string
	Verified     string
	AccessToken  string
	AccessExpire time.Time
	UID          int64 `gorm:"column:uid"`
}

//TableName ..
func (LoginModel) TableName() string {
	return "user_login"
}

//LoginByPassword ..
func (s *LoginModel) LoginByPassword(platform, openid, password string) (*LoginModel, error) {
	pwdHash := md5.New()
	pwdHash.Write([]byte(password))
	pwd := hex.EncodeToString(pwdHash.Sum(nil))
	fmt.Println(pwd)
	var login LoginModel
	return &login, nil
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

// //LoginAdd ..
// func (sess *Session) LoginAdd(uid int64, platform, openid, accessToken string) (*LoginModel, error) {
// 	var login LoginModel
// 	login.UID = uid
// 	login.Openid = openid
// 	login.Platform = platform
// 	login.AccessToken = accessToken
// 	// return &login, nil
// 	result := sess.Create(login)
// 	if result.Error != nil {
// 		return nil, errors.New("账户创建失败")
// 	}
// 	return &login, nil
// }
