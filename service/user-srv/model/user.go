package model

import "time"

//UserModel ..
type UserModel struct {
	UID       int64     `db:"uid,pk"`
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
