package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(&payload)
	if err != nil {
		log.Fatal(err.Error())
	}
}
