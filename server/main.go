package main

import (
	"fmt"

	"github.com/panaka13/travian-server/db"
	"github.com/panaka13/travian-server/model"
)

func create() {
	user := model.User{Name: "panaka_13", Id: 3563}
	fmt.Println(user.Name)
	village := model.Village{Name: "ka", Id: 1234, User: user}
	fmt.Println(village.Name)
	result := db.DB.Create(&user)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	db.DB.Create(&village)
}

func load() {
	var user model.User
	db.DB.First(&user, 3563)
	fmt.Println(user.ID)
	fmt.Println(user.Id)
	var village model.Village
	db.DB.First(&village, 1234)
	fmt.Println(village.ID)
	fmt.Println(village.Id)
}

func main() {
	db.InitDb()
	fmt.Println("Hello World")
	create()
	load()
}
