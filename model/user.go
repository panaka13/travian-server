package model

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id" gorm:"primaryKey"`
}
