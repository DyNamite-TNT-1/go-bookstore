package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //(the blank "_" before this is necessary, read doc of this github for more information)
)

var (
	db *gorm.DB //db is database
)

func Connect() {

	d, err := gorm.Open("mysql", "root:adminnof372001@/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
