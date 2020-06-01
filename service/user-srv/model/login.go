package model

import "time"

//LoginModel ..
type LoginModel struct {
	ID       int64  `db:"id,omitempty"`
	UID      int64  `db:"uid"`
	Platform string `db:"platform"`
	Openid   string `db:"openid"`
	// Verified     string    `db:"verified"`
	AccessToken  string     `db:"access_token"`
	AccessExpire *time.Time `db:"access_expire"`
	Data         string     `db:"data"`
	Status       int64      `db:"status"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
}

//TableName ..
func (LoginModel) TableName() string {
	return "uc_login"
}
