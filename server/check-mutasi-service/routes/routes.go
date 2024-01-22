package routes

import (
	"check-mutasi-service/configs"
	"check-mutasi-service/controllers"
	"check-mutasi-service/repository"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	db := configs.InitDB()

	walletRepository := repository.NewTransactionRepository(db)

	checkMutasiController := controllers.NewCheckMutasi(walletRepository)

	r.HandleFunc("/api/check-mutasi", checkMutasiController.Find).Methods("POST")
}
