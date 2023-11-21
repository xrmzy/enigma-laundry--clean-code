package dto

import "time"

type CustomersRequestDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
