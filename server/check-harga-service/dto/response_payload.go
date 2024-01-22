package dto

type ErrorResponse struct {
	Err     bool   `json:"error"`
	Message string `json:"message"`
}

type Data struct {
	ID           string `json:"id"`
	HargaBuyback int    `json:"harga_buyback"`
	HargaTopup   int    `json:"harga_topup"`
}

type SuccessResponse struct {
	Err  bool `json:"error"`
	Data Data `json:"data"`
}
