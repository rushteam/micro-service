package model

import (
	"errors"
	"log"
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

//LoginByPassword ...
func (sess *Session) LoginByPassword(typ, openid, pwd string) (*LoginModel, error) {
	var login LoginModel
	// return &login, nil
	result := sess.Where("platform = ?", typ).Where("openid = ?", openid).First(&login)
	if result.Error != nil {
		return nil, errors.New("用户不存在")
	}
	log.Println(login.AccessToken, pwd)
	if login.AccessToken != pwd {
		return nil, errors.New("密码错误")
	}
	return &login, nil
}

//LoginAdd ..
func (sess *Session) LoginAdd(uid int64, platform, openid, accessToken string) (*LoginModel, error) {
	var login LoginModel
	login.UID = uid
	login.Openid = openid
	login.Platform = platform
	login.AccessToken = accessToken
	// return &login, nil
	result := sess.Create(login)
	if result.Error != nil {
		return nil, errors.New("账户创建失败")
	}
	return &login, nil
}
