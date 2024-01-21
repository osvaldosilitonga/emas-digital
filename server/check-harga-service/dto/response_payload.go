package dto

type ErrorResponse struct {
	Err     bool   `json:"error"`
	ReffID  string `json:"reff_id"`
	Message string `json:"message"`
}

type Data struct {
	HargaBuyback int `json:"harga_buyback"`
	HargaTopup   int `json:"harga_topup"`
}

type SuccessResponse struct {
	Err  bool `json:"error"`
	Data Data `json:"data"`
}
