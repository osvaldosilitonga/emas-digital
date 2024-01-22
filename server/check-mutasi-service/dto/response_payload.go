package dto

import "check-mutasi-service/entity"

type ErrorResponse struct {
	Err     bool   `json:"error"`
	ReffID  string `json:"reff_id"`
	Message string `json:"message"`
}

type Data struct {
	Norek string  `json:"norek"`
	Saldo float32 `json:"saldo"`
}

type SuccessResponse struct {
	Err  bool                 `json:"error"`
	Data []entity.Transaction `json:"data"`
}
