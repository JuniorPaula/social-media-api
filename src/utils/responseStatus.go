package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	ResponseJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
