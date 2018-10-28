package model

import (
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
}

//TableName ..
func (LoginModel) TableName() string {
	return "user_login"
}

//LoginByPhone ...
func LoginByPhone(phone, pwd string) *LoginModel {
	var login LoginModel
	result := DB.Where("platform = ?", "phone").Where("openid = ?", phone).First(&login)
	if result.Error != nil {
		return nil
	}
	if login.AccessToken != pwd {
		return nil
	}
	return &login
}
