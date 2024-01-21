package controllers

import (
	"check-harga-service/dto"
	"check-harga-service/helpers"
	"net/http"
)

func Find(w http.ResponseWriter, r *http.Request) {
	// find last price from db

	helpers.ResponseJson(w, http.StatusCreated, &dto.SuccessResponse{
		Err: false,
		Data: dto.Data{
			HargaBuyback: 90000,
			HargaTopup:   80000,
		},
	})
}
