package repository

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/mlboy/godb/builder"
	"github.com/mlboy/godb/db"
	// "upper.io/db.v3"
)

var (
	//ErrPassword 密码不正确
	ErrPassword = errors.New(`密码错误`)
	//ErrUnvaildLoginType 无效的登陆方式
	ErrUnvaildLoginType = errors.New("无效的登陆方式")
	//loginMaps 本地登陆方式,即需要加密密码
	loginMaps = map[string]bool{"phone": true, "email": true, "username": true}
)

//Login ..
var Login = newLoginRepo()

//LoginModel ..
type LoginModel struct {
	ID           int64     `db:"id,omitempty"`
	UID          int64     `db:"uid"`
	Platform     string    `db:"platform"`
	Openid       string    `db:"openid"`
	Unionid      string    `db:"unionid"`
	Verified     string    `db:"verified"`
	AccessToken  string    `db:"access_token"`
	AccessExpire time.Time `db:"access_expire"`
	Data         string    `db:"data"`
	Status       int64     `db:"status"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

//TableName ..
func (LoginModel) TableName() string {
	return "uc_login"
}

func newLoginRepo() *loginRepository {
	return &loginRepository{}
}

type loginRepository struct{}

//FindByPassword ...
func (repo loginRepository) FindByPassword(platform, openid, password string) (*LoginModel, error) {
	login := &LoginModel{}
	err := db.Fetch(
		login,
		builder.Where("platform", platform),
		builder.Where("openid", openid),
	)
	if login.AccessToken != password {
		return nil, ErrPassword
	}
	return login, err
}

//Create ..
func (repo loginRepository) Create(login *LoginModel) error {
	if f, ok := loginMaps[login.Platform]; !ok || !f {
		return ErrUnvaildLoginType
	}
	login.Status = 1
	now := time.Now()
	login.CreatedAt = now
	login.UpdatedAt = now
	_, err := db.Insert(login)
	return err
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
