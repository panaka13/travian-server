package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorResponse(err error, w http.ResponseWriter) {
	response := "{\"error\": \"" + err.Error() + "\"}"
	fmt.Fprint(w, response)
}

func ObjectResponse(object interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(object)
	if err != nil {
		ErrorResponse(err, w)
	} else {
		fmt.Fprint(w, json)
	}
}
