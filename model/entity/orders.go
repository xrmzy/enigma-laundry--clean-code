package entity

import "time"

type Orders struct {
	OrderId      string    `json:"orderId"`
	CustomerId   string    `json:"customerId"`
	CustomerName string    `json:"customerName"`
	Service      string    `json:"service"`
	Unit         string    `json:"unit"`
	OutletName   string    `json:"outletName"`
	OrderDate    string    `json:"orderDate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
