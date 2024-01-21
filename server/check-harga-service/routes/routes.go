package routes

import (
	"check-harga-service/configs"
	"check-harga-service/controllers"
	"check-harga-service/repository"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	db := configs.InitDB()

	hargaRepository := repository.NewHargaRepository(db)

	checkHargaController := controllers.NewCheckHarga(hargaRepository)

	r.HandleFunc("/api/check-harga", checkHargaController.FindOne).Methods("GET")
}
