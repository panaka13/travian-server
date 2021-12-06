package main

import (
	"fmt"

	"github.com/panaka13/travian-server/model"
)

func main() {
	fmt.Println("Hello World")
	user := model.User{Name: "panaka_13", Id: 3563}
	fmt.Println(user.Name)
	village := model.Village{Name: "ka", Id: 1234, User: user}
	fmt.Println(village.Name)
}
