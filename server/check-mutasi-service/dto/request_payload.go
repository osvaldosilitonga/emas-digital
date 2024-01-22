package dto

type RequestPayload struct {
	Norek string `json:"norek" validate:"required"`
	Start int64  `json:"start_date" validate:"required"`
	End   int64  `json:"end_date" validate:"required"`
}
