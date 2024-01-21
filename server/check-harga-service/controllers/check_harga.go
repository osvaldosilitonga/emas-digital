package controllers

import (
	"check-harga-service/dto"
	"check-harga-service/helpers"
	"check-harga-service/repository"
	"context"
	"log"
	"net/http"
	"time"
)

type CheckHarga struct {
	HargaRepo *repository.HargaRepository
}

func NewCheckHarga(hr *repository.HargaRepository) *CheckHarga {
	return &CheckHarga{
		HargaRepo: hr,
	}
}

func (ch *CheckHarga) FindOne(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, err := ch.HargaRepo.FindOne(ctx)
	if err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
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
