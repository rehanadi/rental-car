package controllers

import (
	"net/http"
	"rental-car/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CarController struct {
	CarRepository repositories.CarRepository
}

func NewCarController(cr repositories.CarRepository) *CarController {
	return &CarController{cr}
}

func (cc *CarController) GetAllCars(c echo.Context) error {
	cars, statusCode, err := cc.CarRepository.FindAll()
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get all cars", "cars": cars})
}

func (cc *CarController) GetCarById(c echo.Context) error {
	carId := c.Param("id")
	id, err := strconv.Atoi(carId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid car id"})
	}

	car, statusCode, err := cc.CarRepository.FindById(id)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get car by id", "car": car})
}

func (cc *CarController) GetCarsByCategoryId(c echo.Context) error {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid category id"})
	}

	cars, statusCode, err := cc.CarRepository.FindByCategoryId(id)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get cars by category id", "cars": cars})
}
