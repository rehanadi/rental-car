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

	paymentRepository := repositories.NewPaymentRepository(config.DB)
	paymentController := controllers.NewPaymentController(paymentRepository)

	categoryRepository := repositories.NewCategoryRepository(config.DB)
	categoryController := controllers.NewCategoryController(categoryRepository)

	carRepository := repositories.NewCarRepository(config.DB)
	carController := controllers.NewCarController(carRepository)

	rentalRepository := repositories.NewRentalRepository(config.DB)
	rentalController := controllers.NewRentalController(rentalRepository)

	u := e.Group("/users")
	u.POST("/register", userController.Register)
	u.POST("/login", userController.Login)
	u.GET("/me", m.Authentication(userController.GetUserProfile))

	p := e.Group("/payments")
	p.POST("/top-up", m.Authentication(paymentController.TopUpDeposit))
	p.GET("/verify/:id", paymentController.VerifyPayment)

	c := e.Group("/categories")
	c.Use(m.Authentication)
	c.GET("", categoryController.GetAllCategories)

	car := e.Group("/cars")
	car.Use(m.Authentication)
	car.GET("", carController.GetAllCars)
	car.GET("/:id", carController.GetCarById)
	car.GET("/category/:id", carController.GetCarsByCategoryId)
	car.POST("/rent", rentalController.RentCar)
	car.POST("/return/:id", rentalController.ReturnCar)
}
