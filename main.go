package main

import (
	"os"
	"rental-car/config"
	"rental-car/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitDB()
	defer config.CloseDB()

	e := echo.New()
	routes.Init(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
