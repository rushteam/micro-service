package model

import (
	"errors"
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
func LoginByPassword(typ, openid, pwd string) (*LoginModel, error) {
	var login LoginModel
	// return &login, nil
	result := Db().Where("platform = ?", typ).Where("openid = ?", openid).First(&login)
	if result.Error != nil {
		return nil, errors.New("用户不存在")
	}
	if login.AccessToken != pwd {
		return nil, errors.New("密码错误")
	}
	return &login, nil
}
