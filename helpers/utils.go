package helpers

func CalculateCost(rentalCost float64, rentalDays int) (subtotalCost, taxCost, totalCost float64) {
	subtotalCost = rentalCost * float64(rentalDays)
	taxCost = subtotalCost * 0.12
	totalCost = subtotalCost + taxCost
	return
}
