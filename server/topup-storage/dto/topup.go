package dto

type Topup struct {
	ReffID string  `json:"reff_id"`
	Gram   float32 `json:"gram"`
	Harga  int     `json:"harga"`
	NoRek  string  `json:"norek"`
	Saldo  float32 `json:"balance"`
}
