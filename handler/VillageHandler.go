package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/panaka13/travian-server/db"
	"github.com/panaka13/travian-server/model"
)

func CreateVillageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	var request map[string]interface{}
	json.NewDecoder(r.Body).Decode(&request)
	if err := CheckBodyParams(request, "id", "name", "user"); err != nil {
		ErrorResponse(err, w)
		fmt.Println(err)
		return
	}

	// parse village data
	var village model.Village
	village.Id = int(request["id"].(float64))
	village.Name = request["name"].(string)
	db.DB.First(&village.User, int(request["user"].(float64)))

	result := db.DB.Create(&village)
	db.DB.Save(&village)
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
	} else {
		ObjectResponse(village, w)
	}
}

func GetVillageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, _ := vars["villageid"]
	var village model.Village
	result := db.DB.Joins("User").First(&village, id)
	village.Deserialize()
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
	} else {
		ObjectResponse(village, w)
	}
}

func UpdateVillage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		HandlePreflight(w, r)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	var request map[string]interface{}
	json.NewDecoder(r.Body).Decode(&request)
	if err := CheckBodyParams(request, "villageid", "part", "structure"); err != nil {
		ErrorResponse(err, w)
		fmt.Println(err)
		return
	}

	var village model.Village
	result := db.DB.First(&village, int(request["villageid"].(float64)))
	if result.Error != nil {
		ErrorResponse(result.Error, w)
		fmt.Println(result.Error)
		return
	}

	village.Deserialize()
	wood := int(request["wood"].(float64))
	clay := int(request["clay"].(float64))
	iron := int(request["iron"].(float64))
	wheat := int(request["wheat"].(float64))
	village.UpdateResource(wood, clay, iron, wheat)
	village.PartialDeserialize(request["structure"].(string), request["part"].(string))
	village.Serialize()

	db.DB.Save(village)
	ObjectResponse(village, w)
}
