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
		fmt.Fprint(w, string(json))
	}
}

func CheckBodyParams(params map[string]interface{}, keys ...string) error {
	for _, key := range keys {
		if _, exist := params[key]; !exist {
			return fmt.Errorf("Body missing key %s", key)
		}
	}
	return nil
}
