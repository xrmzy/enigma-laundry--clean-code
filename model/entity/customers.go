package entity

import "time"

type Customer struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
