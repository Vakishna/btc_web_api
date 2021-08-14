package database

import (
	"github.com/Vakishna/WebWalApiAuth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:.k@/lightningapp"), &gorm.Config{})

	if err != nil {
		panic("Database not reachable!")
	}
	DB = db
	db.AutoMigrate(&models.User{})

}
