package routes

import (
	"check-harga-service/configs"
	"check-harga-service/controllers"
	"log"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	db := configs.InitDB()
	if err := db.Ping(); err != nil {
		log.Println("DB Connection Fail")
	}

	r.HandleFunc("/api/check-harga", controllers.Find).Methods("GET")
}
