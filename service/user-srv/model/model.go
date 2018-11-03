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

//Db ..
func Db() *gorm.DB {
	db, err := pool.GetDb(db.Default, db.Master)
	if err != nil {
		panic(err.Error())
	}
	return db
}
