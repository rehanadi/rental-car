package models

type Category struct {
	CategoryId  int    `json:"category_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
