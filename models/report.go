package models

import "time"

type ReportRentalDetail struct {
	RentalID     int       `json:"rental_id"`
	CarID        int       `json:"car_id"`
	CarName      string    `json:"car_name"`
	RentalCost   float64   `json:"rental_cost"`
	RentalDays   int       `json:"rental_days"`
	SubtotalCost float64   `json:"subtotal_cost"`
	TaxCost      float64   `json:"tax_cost"`
	TotalCost    float64   `json:"total_cost"`
	Status       string    `json:"status"`
	RentedAt     time.Time `json:"rented_at"`
	ExpiredAt    time.Time `json:"expired_at"`
	ReturnedAt   time.Time `json:"returned_at"`
}

type ReportRentalSummary struct {
	TotalCar    int     `json:"total_car"`
	TotalRental int     `json:"total_rental"`
	TotalCost   float64 `json:"total_cost"`
}
