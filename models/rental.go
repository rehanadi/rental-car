package models

import "time"

type Rental struct {
	RentalID     int       `json:"rental_id" gorm:"primaryKey"`
	UserID       int       `json:"user_id"`
	CarID        int       `json:"car_id"`
	RentalCost   float64   `json:"rental_cost"`
	RentalDays   int       `json:"rental_days"`
	SubtotalCost float64   `json:"subtotal_cost"`
	TaxCost      float64   `json:"tax_cost"`
	TotalCost    float64   `json:"total_cost"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiredAt    time.Time `json:"expired_at"`
	ReturnedAt   time.Time `json:"returned_at"`
}
