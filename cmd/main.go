package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	app := echo.New()

	port := os.Getenv("PORT")

	app.Start(":" + port)
}
