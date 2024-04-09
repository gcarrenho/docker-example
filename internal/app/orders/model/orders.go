package model

import "time"

type OrdersResponse struct {
	OrderNumber string    `json:"order_number"`
	Currency    string    `json:"currency"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"creates_at"`
}
