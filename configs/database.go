package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(mysql.Open("root:elder@tcp(localhost:1010)/gin-api?charset=tru"), &gorm.Config{})

	if err != nil {
		panic("Failed connect db")
	}

	DB = db
}
