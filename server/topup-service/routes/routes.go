package routes

import (
	"topup-service/controllers"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/topup", controllers.Add).Methods("POST")
}
