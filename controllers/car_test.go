package controllers

import (
	"rental-car/models"
	"rental-car/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var carRepoMock = &repositories.CarRepoMock{Mock: mock.Mock{}}
var carController = CarMockController{CarRepository: carRepoMock}

func TestGetCarNotFound(t *testing.T) {
	var id int = 1
	carRepoMock.Mock.On("FindById", id).Return(nil)
	car, err := carController.FindById(id)

	assert.Nil(t, car)
	assert.NotNil(t, err)
	assert.Equal(t, "car not found", err.Error(), "error message should be 'car not found'")
}

func TestGetCarFound(t *testing.T) {
	var id int = 2
	resCar := models.Car{CarID: 2, Name: "Car 2", CategoryID: 1}

	carRepoMock.Mock.On("FindById", id).Return(resCar)
	car, err := carController.FindById(id)

	assert.Nil(t, err)
	assert.NotNil(t, car)
}
