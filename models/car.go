package models

type Car struct {
	CarID       int     `json:"car_id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	RentalCost  float64 `json:"rental_cost"`
	Stock       int     `json:"stock"`
	CategoryID  int     `json:"category_id"`
}

type CarResponse struct {
	CarID        int     `json:"car_id" gorm:"primaryKey"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	RentalCost   float64 `json:"rental_cost"`
	Stock        int     `json:"stock"`
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
}
