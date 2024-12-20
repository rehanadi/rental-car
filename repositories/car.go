package repositories

import (
	"errors"
	"net/http"
	"rental-car/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	FindAll() ([]models.Car, int, error)
	FindById(id int) (models.Car, int, error)
	FindByCategoryId(categoryId int) ([]models.Car, int, error)
}

type carRepository struct {
	DB *gorm.DB
}

func NewCarRepository(DB *gorm.DB) *carRepository {
	return &carRepository{DB}
}

func (r *carRepository) FindAll() ([]models.Car, int, error) {
	var cars []models.Car
	if err := r.DB.Table("cars").Select("cars.*, categories.name as category_name").Joins("left join categories on cars.category_id = categories.category_id").Find(&cars).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return cars, http.StatusOK, nil
}

func (r *carRepository) FindById(id int) (models.Car, int, error) {
	var car models.Car
	if err := r.DB.Table("cars").Select("cars.*, categories.name as category_name").Joins("left join categories on cars.category_id = categories.category_id").Where("car_id = ?", id).First(&car).Error; err != nil {
		return car, http.StatusNotFound, errors.New("car not found")
	}

	return car, http.StatusOK, nil
}

func (r *carRepository) FindByCategoryId(categoryId int) ([]models.Car, int, error) {
	var cars []models.Car
	if err := r.DB.Table("cars").Select("cars.*, categories.name as category_name").Joins("left join categories on cars.category_id = categories.category_id").Where("cars.category_id = ?", categoryId).Find(&cars).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return cars, http.StatusOK, nil
}
