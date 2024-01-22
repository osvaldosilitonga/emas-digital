package entity

type Transaction struct {
	Date         int64   `json:"date"`
	Type         string  `json:"type"`
	Gram         float32 `json:"gram"`
	HargaTopup   int     `json:"harga_topup"`
	HargaBuyback int     `json:"harga_buyback"`
	Balance      float32 `json:"balance"`
}

// type Transaction struct {
// 	ID           int     `json:"id"`
// 	WalletID     string  `json:"wallet_id"`
// 	Date         int64   `json:"date"`
// 	Type         string  `json:"type"`
// 	Gram         float32 `json:"gram"`
// 	Price        int     `json:"price"`
// 	HargaTopup   int     `json:"harga_topup"`
// 	HargaBuyback int     `json:"harga_buyback"`
// 	Balance      float32 `json:"balance"`
// }
