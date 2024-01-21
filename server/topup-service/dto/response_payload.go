package dto

type ErrorResponse struct {
	Err     bool   `json:"error"`
	ReffID  string `json:"reff_id"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Err    bool   `json:"error"`
	ReffID string `json:"reff_id"`
}
