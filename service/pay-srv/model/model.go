package model

import (
	"gitee.com/rushteam/micro-service/common/db"
	"github.com/jinzhu/gorm"
)

//DbPool ..
type DbPool interface {
	GetDb(groupName string, pool string) (*gorm.DB, error)
}

//pool ..
var pool DbPool

//Init ..
func Init(dbPool DbPool) {
	pool = dbPool
}

//Session ..
type Session struct {
	*gorm.DB
}

//Db ..
func Db() *Session {
	db, err := pool.GetDb(db.Default, db.Master)
	if err != nil {
		panic(err.Error())
	}
	return &Session{db}
}

//Begin ..
func Begin() *Session {
	return &Session{Db().Begin()}
}
