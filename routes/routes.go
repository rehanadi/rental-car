package routes

import (
	"rental-car/config"
	"rental-car/controllers"
	m "rental-car/middlewares"
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

	// payment routes
	paymentRepository := repositories.NewPaymentRepository(config.DB)
	paymentController := controllers.NewPaymentController(paymentRepository)

	p := e.Group("/payments")
	p.POST("/top-up", m.Authentication(paymentController.TopUpDeposit))
	p.GET("/verify/:id", paymentController.VerifyPayment)
}
