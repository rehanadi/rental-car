package repositories

import (
	"errors"
	"fmt"
	"net/http"
	"rental-car/helpers"
	"rental-car/models"
	"time"

	"gorm.io/gorm"
)

type RentalRepository interface {
	RentCar(rentCarRequest *models.RentCarRequest) (*models.Rental, int, error)
	ReturnCar(rentalID int) (int, error)
}

type rentalRepository struct {
	DB *gorm.DB
}

func NewRentalRepository(DB *gorm.DB) *rentalRepository {
	return &rentalRepository{DB}
}

func (r *rentalRepository) RentCar(rentCarRequest *models.RentCarRequest) (*models.Rental, int, error) {
	// check if user exists
	var user models.User
	if err := r.DB.Where("user_id = ?", rentCarRequest.UserID).
		First(&user).Error; err != nil {
		return nil, http.StatusNotFound, errors.New("user not found")
	}

	// check if car exists
	var car models.Car
	if err := r.DB.Where("car_id = ?", rentCarRequest.CarID).
		First(&car).Error; err != nil {
		return nil, http.StatusNotFound, errors.New("car not found")
	}

	// check if car stock is enough
	if car.Stock <= 0 {
		return nil, http.StatusBadRequest, errors.New("car out of stock")
	}

	// calculate cost
	subtotalCost, taxCost, totalCost := helpers.CalculateCost(car.RentalCost, rentCarRequest.RentalDays)

	// check if user has enough deposit
	if user.DepositAmount < totalCost {
		return nil, http.StatusBadRequest, errors.New("insufficient deposit Rp " + fmt.Sprintf("%.2f", user.DepositAmount) + " for total cost Rp " + fmt.Sprintf("%.2f", totalCost))
	}

	// deduct user deposit
	user.DepositAmount -= totalCost
	err := r.DB.Save(&user).Error
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// deduct car stock
	car.Stock--
	err = r.DB.Save(&car).Error
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// insert into rentals
	rental := models.Rental{
		UserID:       rentCarRequest.UserID,
		CarID:        rentCarRequest.CarID,
		RentalCost:   car.RentalCost,
		RentalDays:   rentCarRequest.RentalDays,
		SubtotalCost: subtotalCost,
		TaxCost:      taxCost,
		TotalCost:    totalCost,
		Status:       "ongoing",
		RentedAt:     time.Now(),
		ExpiredAt:    time.Now().AddDate(0, 0, rentCarRequest.RentalDays),
	}

	result := r.DB.Create(&rental)
	if result.Error != nil {
		return nil, http.StatusInternalServerError, result.Error
	}

	return &rental, http.StatusOK, nil
}

func (r *rentalRepository) ReturnCar(rentalID int) (int, error) {
	// check if rental exists
	var rental models.Rental
	if err := r.DB.Where("rental_id = ?", rentalID).
		First(&rental).Error; err != nil {
		return http.StatusNotFound, errors.New("rental not found")
	}

	// check if car exists
	var car models.Car
	if err := r.DB.Where("car_id = ?", rental.CarID).
		First(&car).Error; err != nil {
		return http.StatusNotFound, errors.New("car not found")
	}

	// check if user exists
	var user models.User
	if err := r.DB.Where("user_id = ?", rental.UserID).
		First(&user).Error; err != nil {
		return http.StatusNotFound, errors.New("user not found")
	}

	// check if rental status is ongoing
	if rental.Status != "ongoing" {
		return http.StatusBadRequest, errors.New("rental already returned")
	}

	// calculate cost penalty if rental is late using helpers.CalculateCost
	if time.Now().After(rental.ExpiredAt) {
		lateRentalDays := int(time.Since(rental.ExpiredAt).Hours() / 24)
		subtotalCost, taxCost, totalCost := helpers.CalculateCost(car.RentalCost, lateRentalDays)

		// check if user has enough deposit
		if user.DepositAmount < totalCost {
			return http.StatusBadRequest, errors.New("insufficient deposit Rp " + fmt.Sprintf("%.2f", user.DepositAmount) + " for total cost Rp " + fmt.Sprintf("%.2f", totalCost))
		}

		// deduct user deposit
		user.DepositAmount -= totalCost
		err := r.DB.Save(&user).Error
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// update rental cost
		rental.RentalDays += lateRentalDays
		rental.SubtotalCost += subtotalCost
		rental.TaxCost += taxCost
		rental.TotalCost += totalCost

		err = r.DB.Save(&rental).Error
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	// return car stock
	car.Stock++
	err := r.DB.Save(&car).Error
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// update rental status
	rental.Status = "finished"
	rental.ReturnedAt = time.Now()
	err = r.DB.Save(&rental).Error
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
