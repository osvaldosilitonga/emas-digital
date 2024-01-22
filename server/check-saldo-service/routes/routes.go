package routes

import (
	"check-saldo-service/configs"
	"check-saldo-service/controllers"
	"check-saldo-service/repository"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	db := configs.InitDB()

	walletRepository := repository.NewWalletRepository(db)

	checkSaldoController := controllers.NewCheckSaldo(walletRepository)

	r.HandleFunc("/api/check-saldo", checkSaldoController.FindOne).Methods("POST")
}
