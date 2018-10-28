package db

import (
	"github.com/go-log/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//InitDB 初始化db数据
func InitDB(sourceDSN string) *gorm.DB {
	var err error
	DB, err = gorm.Open("mysql", sourceDSN)
	// defer db.Close()
	if err != nil {
		log.Log("[db] open fail (%s) %s", sourceDSN, err)
	}
	return DB
}
