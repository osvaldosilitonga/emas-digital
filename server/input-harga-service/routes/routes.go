package routes

import (
	"input-harga-service/controllers"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/input-harga", controllers.Add).Methods("POST")
}
