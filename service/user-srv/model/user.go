package model

import (
	"errors"
	"time"
)

//UserModel ..
type UserModel struct {
	// gorm.Model
	UID       int64 `gorm:"column:uid;PRIMARY_KEY;AUTO_INCREMENT;"`
	Nickname  string
	Gender    string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//TableName ..
func (UserModel) TableName() string {
	return "user_user"
}

//UserByUID ...
func (UserModel) UserByUID(uid int64) (*UserModel, error) {
	var user UserModel
	result := DB.Where("uid = ?", uid).First(&user)
	if result.Error != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, nil
}
