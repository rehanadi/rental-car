package controllers

import (
	"errors"
	"rental-car/models"
	"rental-car/repositories"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryRepository repositories.CategoryRepository
}

type CategoryMockController struct {
	CategoryRepository repositories.CategoryMockRepository
}

func NewCategoryController(cr repositories.CategoryRepository) *CategoryController {
	return &CategoryController{cr}
}

func (cc *CategoryController) GetAllCategories(c echo.Context) error {
	categories, statusCode, err := cc.CategoryRepository.FindAll()
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get all categories", "categories": categories})
}

func (cmc CategoryMockController) FindAll() (*[]models.Category, error) {
	cars := cmc.CategoryRepository.FindAll()

	if cars == nil {
		return nil, errors.New("category not found")
	}

	return cars, nil
}
