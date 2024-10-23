package model

type Order struct {
	OrderID      string `json:"orderId"`
	CustomerName string `json:"customerName"`
	Product      string `json:"product"`
	Quantity     int    `json:"quantity"`
}
