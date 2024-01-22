package controllers

import (
	"buyback-service/dto"
	"buyback-service/handlers"
	"buyback-service/helpers"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/teris-io/shortid"
)

var (
	validate *validator.Validate = validator.New()
)

func Add(w http.ResponseWriter, r *http.Request) {
	sid, _ := shortid.Generate()

	payload := dto.RequestPayload{}
	payload.ReffID = sid

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: err.Error(),
		})
		log.Println(err.Error())
		return
	}
	defer r.Body.Close()

	if err := validate.Struct(&payload); err != nil {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: err.Error(),
		})
		log.Println(err.Error())
		return
	}

	// Check digit gram
	gramStr := fmt.Sprintf("%v", payload.Gram)
	dotIndex := strings.Index(gramStr, ".")
	if dotIndex != -1 {
		decimalDigits := len(gramStr) - dotIndex - 1
		if decimalDigits > 3 {
			helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
				Err:     true,
				ReffID:  payload.ReffID,
				Message: "maksimal 3 digit dibelakang koma",
			})

			return
		}
	}

	// Check Harga
	res, id, err := handlers.CheckHarga(payload.Harga)
	if err != nil {
		helpers.ResponseJson(w, http.StatusInternalServerError, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: err.Error(),
		})
		log.Println(err.Error())
		return
	}
	if !res {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: "harga tidak sesuai dengan harga saat ini",
		})
		return
	}

	// Check Saldo
	rs, err := handlers.CheckSaldo(payload.NoRek, payload.Gram)
	if err != nil {
		helpers.ResponseJson(w, http.StatusInternalServerError, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: err.Error(),
		})
		log.Println(err.Error())
		return
	}
	if !rs {
		helpers.ResponseJson(w, http.StatusBadRequest, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: "saldo tidak cukup",
		})
		return
	}

	payload.PriceID = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := handlers.BuybackPublish(ctx, &payload); err != nil {
		helpers.ResponseJson(w, http.StatusInternalServerError, &dto.ErrorResponse{
			Err:     true,
			ReffID:  payload.ReffID,
			Message: "Kafka not ready",
		})
		return
	}

	helpers.ResponseJson(w, http.StatusCreated, &dto.SuccessResponse{
		Err:    false,
		ReffID: payload.ReffID,
	})
}
