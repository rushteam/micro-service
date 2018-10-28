package db

import (
	"log"
	"time"

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
		log.Printf("[db] connect failed (%s) %s\r\n", sourceDSN, err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	DB.DB().SetConnMaxLifetime(time.Hour)
	return DB
}
