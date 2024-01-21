package dto

type RequestPayload struct {
	SID          string `json:"sid,omitempty"`
	AdminID      string `json:"admin_id" validate:"required"`
	HargaTopup   int    `json:"harga_topup" validate:"required,min=1000"`
	HargaBuyback int    `json:"harga_buyback" validate:"required,min=1000"`
}
