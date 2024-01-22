package controllers

import (
	"check-mutasi-service/dto"
	"check-mutasi-service/helpers"
	"check-mutasi-service/repository"
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

type CheckMutasi struct {
	TransactionRepo *repository.TransactionRepository
}

func NewCheckMutasi(tr *repository.TransactionRepository) *CheckMutasi {
	return &CheckMutasi{
		TransactionRepo: tr,
	}
}

func (cm *CheckMutasi) Find(w http.ResponseWriter, r *http.Request) {
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

	data, err := cm.TransactionRepo.FindByNorek(ctx, payload.Norek, payload.Start, payload.End)
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
		Err:  false,
		Data: data,
	})
}
