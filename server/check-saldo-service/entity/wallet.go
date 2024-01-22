package entity

type Wallets struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Balance    float32 `json:"balance"`
}
