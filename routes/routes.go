package routes

import (
	"rental-car/config"
	"rental-car/controllers"
	"rental-car/repositories"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	// user routes
	userRepository := repositories.NewUserRepository(config.DB)
	userController := controllers.NewUserController(userRepository)

	u := e.Group("/users")
	u.POST("/register", userController.Register)
	u.POST("/login", userController.Login)
}
