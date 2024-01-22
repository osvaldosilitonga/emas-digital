package dto

type Data struct {
	ID         string `json:"id"`
	HargaTopup int    `json:"harga_topup"`
}

type APIResponse struct {
	Data Data `json:"data"`
}
