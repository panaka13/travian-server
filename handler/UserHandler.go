package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/panaka13/travian-server/db"
	"github.com/panaka13/travian-server/model"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	result := db.DB.Create(&user)
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
	} else {
		ObjectResponse(user, w)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
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
