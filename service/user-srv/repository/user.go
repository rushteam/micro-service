package repository

import (
	"time"

	"upper.io/db.v3"
)

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

//UserRepository ..
type UserRepository struct {
	Db Database
}

//FindByUID ...
func (repo UserRepository) FindByUID(uid int64) (*UserModel, error) {
	var user UserModel
	res := repo.Db.Collection(user.TableName()).Find(db.Cond{
		"uid": uid,
	})
	// fmt.Println(res)
	err := res.One(&user)
	return &user, err
}

//Create ..
func (repo UserRepository) Create(user *UserModel) (int64, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	uid, err := repo.Db.Collection(user.TableName()).Insert(user)
	user.UID = uid.(int64)
	return uid.(int64), err
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
