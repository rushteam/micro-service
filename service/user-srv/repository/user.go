package repository

import (
	"time"

	"github.com/mlboy/godb/builder"
	"github.com/mlboy/godb/db"
)

//User ..
var User = newUserRepo()

//UserModel ..
type UserModel struct {
	UID       int64     `db:"uid"`
	Nickname  string    `db:"nickname"`
	Gender    string    `db:"gender"`
	Avatar    string    `db:"avatar"`
	Status    int32     `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

//TableName ..
func (UserModel) TableName() string {
	return "uc_user"
}

func newUserRepo() *userRepository {
	return &userRepository{}
}

//userRepository ..
type userRepository struct{}

//FindByUID ...
func (repo userRepository) FindByUID(uid int64) (*UserModel, error) {
	user := &UserModel{}
	err := db.Fetch(
		user,
		builder.Where("uid", uid),
	)
	return user, err
}

//Create ..
func (repo userRepository) Create(user *UserModel) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := db.Insert(user)
	return err
}

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
