package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/panaka13/travian-server/db"
	"github.com/panaka13/travian-server/model"
)

func FindUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	name := r.URL.Query().Get("name")

	var user model.User
	result := db.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
	} else {
		ObjectResponse(user, w)
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	var request map[string]interface{}
	json.NewDecoder(r.Body).Decode(&request)
	if err := CheckBodyParams(request, "id", "name"); err != nil {
		ErrorResponse(err, w)
		fmt.Println(err)
		return
	}

	var user model.User
	user.Id = int(request["id"].(float64))
	user.Name = request["name"].(string)

	fmt.Println(user)
	result := db.DB.Create(&user)
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
	} else {
		ObjectResponse(user, w)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, _ := vars["userid"]
	var user model.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
	} else {
		ObjectResponse(user, w)
	}
}
