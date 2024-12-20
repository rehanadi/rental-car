package repositories

import (
	"net/http"
	"rental-car/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, int, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) *categoryRepository {
	return &categoryRepository{DB}
}

func (r *categoryRepository) FindAll() ([]models.Category, int, error) {
	var categories []models.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return categories, http.StatusOK, nil
}
