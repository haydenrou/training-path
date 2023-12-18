package main

import (
	"os"

	"github.com/haydenrou/training-path/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	app := echo.New()

	port := os.Getenv("PORT")

	homeHandler := handler.HomeHandler{}
	app.GET("/", homeHandler.HandleHomeShow)

	app.Start(":" + port)
}
