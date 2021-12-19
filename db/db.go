package db

import (
	"fmt"

	"github.com/panaka13/travian-server/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	SQLITE3  = "sqlite3"
	POSTGRES = "postgres"
)

var (
	DB *gorm.DB
)

func InitDb(databaseType string, databaseUrl string) {
	var err error
	if databaseType == SQLITE3 {
		DB, err = gorm.Open(sqlite.Open(databaseUrl), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if databaseType == POSTGRES {
		DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		fmt.Println("Cannot connect database")
		DB = nil
		return
	}
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Village{})
}
