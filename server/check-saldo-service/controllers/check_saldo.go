package controllers

import (
	"check-saldo-service/dto"
	"check-saldo-service/helpers"
	"check-saldo-service/repository"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/teris-io/shortid"
)

var (
	validate *validator.Validate = validator.New()
)

type CheckSaldo struct {
	WalletRepo *repository.WalletRepository
}

func NewCheckSaldo(wr *repository.WalletRepository) *CheckSaldo {
	return &CheckSaldo{
		WalletRepo: wr,
	}
}

func (ch *CheckSaldo) FindOne(w http.ResponseWriter, r *http.Request) {
	sid, _ := shortid.Generate()

	payload := dto.RequestPayload{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  sid,
			Message: err.Error(),
		})

		log.Println(err.Error())
		return
	}
	defer r.Body.Close()

	if err := validate.Struct(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  sid,
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, err := ch.WalletRepo.FindByNorek(ctx, payload.Norek)
	if err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  sid,
			Message: err.Error(),
		})

		log.Println(err.Error())
		return
	}

	helpers.ResponseJson(w, http.StatusCreated, &dto.SuccessResponse{
		Err: false,
		Data: dto.Data{
			Norek: data.ID,
			Saldo: data.Balance,
		},
	})
}
