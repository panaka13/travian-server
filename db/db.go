package db

import (
	"fmt"

	"github.com/panaka13/travian-server/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect database")
		DB = nil
		return
	}
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Village{})
}
