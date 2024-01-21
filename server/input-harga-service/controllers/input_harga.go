package controllers

import (
	"encoding/json"
	"input-harga-service/dto"
	"input-harga-service/handlers"
	"input-harga-service/helpers"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/teris-io/shortid"
)

var (
	validate *validator.Validate = validator.New()
)

func Add(w http.ResponseWriter, r *http.Request) {
	sid, _ := shortid.Generate()

	payload := dto.RequestPayload{}
	payload.SID = sid

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.SID,
			Message: err.Error(),
		})
		log.Println(err.Error())
		return
	}
	defer r.Body.Close()

	if err := validate.Struct(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.SID,
			Message: err.Error(),
		})
		log.Println(err.Error())
		return
	}

	if err := handlers.HargaPublish(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusInternalServerError, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.SID,
			Message: "Kafka not ready",
		})
		return
	}

	helpers.ResponseJson(w, http.StatusCreated, &dto.SuccessResponse{
		Err:    false,
		ReffID: payload.SID,
	})
}
