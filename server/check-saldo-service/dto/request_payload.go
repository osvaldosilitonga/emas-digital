package dto

type RequestPayload struct {
	Norek string `json:"norek" validate:"required"`
}
