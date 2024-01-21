package dto

type RequestPayload struct {
	ReffID string  `json:"reff_id,omitempty"`
	Gram   float32 `json:"gram" validate:"required,min=0.001"`
	Harga  int     `json:"harga" validate:"required,min=1000"`
	NoRek  string  `json:"norek" validate:"required"`
}
