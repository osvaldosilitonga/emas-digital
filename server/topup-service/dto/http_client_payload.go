package dto

type Data struct {
	HargaTopup int `json:"harga_topup"`
}

type APIResponse struct {
	Data Data `json:"data"`
}
