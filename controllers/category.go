package controllers

import (
	"rental-car/repositories"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryRepository repositories.CategoryRepository
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
