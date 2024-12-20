package controllers

import (
	"net/http"
	"rental-car/models"
	"rental-car/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RentalController struct {
	RentalRepository repositories.RentalRepository
}

func NewRentalController(rr repositories.RentalRepository) *RentalController {
	return &RentalController{rr}
}

func (rc *RentalController) RentCar(c echo.Context) error {
	var rentCarRequest models.RentCarRequest
	c.Bind(&rentCarRequest)

	userId := c.Get("user_id").(int)
	rentCarRequest.UserID = userId

	if rentCarRequest.UserID == 0 || rentCarRequest.CarID == 0 || rentCarRequest.RentalDays == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "user_id, car_id, and rental_days are required"})
	}

	if rentCarRequest.RentalDays < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "rental_days must be greater than 0"})
	}

	rental, statusCode, err := rc.RentalRepository.RentCar(&rentCarRequest)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success rent car", "rental": rental})
}

func (rc *RentalController) ReturnCar(c echo.Context) error {
	rentalId := c.Param("id")
	id, err := strconv.Atoi(rentalId)
	if err != nil || id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid rental id"})
	}

	statusCode, err := rc.RentalRepository.ReturnCar(id)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]string{"message": "success return car"})
}
