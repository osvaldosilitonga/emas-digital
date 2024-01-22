package routes

import (
	"buyback-service/controllers"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/buyback", controllers.Add).Methods("POST")
}
