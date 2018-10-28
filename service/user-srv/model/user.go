package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Login struct {
	gorm.Model
	ID           int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Platform     string
	Openid       string
	Verified     string
	AccessToken  string
	AccessExpire time.Time
}

func (Login) TableName() string {
	return "user_login"
}

func UserLogin(db) {
	db.Where("openid = ?", "jinzhu").First(&user)
}
