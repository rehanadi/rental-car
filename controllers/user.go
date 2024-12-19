package controllers

import (
	"rental-car/models"
	"rental-car/repositories"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserRepository repositories.UserRepository
}

func NewUserController(ur repositories.UserRepository) *UserController {
	return &UserController{ur}
}

func (uc *UserController) Register(c echo.Context) error {
	var user models.RegisterRequest
	c.Bind(&user)

	registerResponse, statusCode, err := uc.UserRepository.Register(&user)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success register", "user": registerResponse})
}

func (uc *UserController) Login(c echo.Context) error {
	var user models.LoginRequest
	c.Bind(&user)

	token, statusCode, err := uc.UserRepository.Login(&user)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success login", "token": token})
}

func (uc *UserController) GetUserProfile(c echo.Context) error {
	userId := c.Get("user_id").(int)

	user, statusCode, err := uc.UserRepository.GetUserProfile(userId)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get user profile", "user": user})
}
