package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Id   int    `json:"id" gorm:"primaryKey"`
}
