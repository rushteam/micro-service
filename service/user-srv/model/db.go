package model

import "github.com/jinzhu/gorm"

var (
	//DB ..
	DB *gorm.DB
)

//SetDB ..
func SetDB(db *gorm.DB) {
	DB = db
}
