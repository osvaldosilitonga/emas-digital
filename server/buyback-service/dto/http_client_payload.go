package dto

type Data struct {
	ID           string `json:"id"`
	HargaTopup   int    `json:"harga_topup"`
	HargaBuyback int    `json:"harga_buyback"`
}

type APIResponse struct {
	Data Data `json:"data"`
}

type CheckSaldoResponse struct {
	Saldo DataSaldo `json:"data"`
}

type DataSaldo struct {
	Norek string  `json:"norek"`
	Saldo float32 `json:"saldo"`
}
