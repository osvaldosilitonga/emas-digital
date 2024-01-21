package dto

type Harga struct {
	ID           string `json:"sid"`
	AdminID      string `json:"admin_id"`
	HargaTopup   int    `json:"harga_topup"`
	HargaBuyback int    `json:"harga_buyback"`
}
