package controllers

import (
	"rental-car/models"
	"rental-car/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepoMock = &repositories.CategoryRepoMock{Mock: mock.Mock{}}
var categoryController = CategoryMockController{CategoryRepository: categoryRepoMock}

func TestGetCategoryNotFound(t *testing.T) {
	categoryRepoMock.Mock.On("FindAll").Return(nil)
	categories, err := categoryController.FindAll()

	assert.Nil(t, categories)
	assert.NotNil(t, err)
	assert.Equal(t, "category not found", err.Error(), "error message should be 'category not found'")
}

func TestGetCategoryFound(t *testing.T) {
	resCategories := []models.Category{
		{CategoryId: 1, Name: "Category 1"},
		{CategoryId: 2, Name: "Category 2"},
	}

	categoryRepoMock.Mock.On("FindAll").Return(resCategories)
	categories, err := categoryController.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, categories)
}
