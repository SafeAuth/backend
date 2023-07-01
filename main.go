package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/SafeAuth/backend/database"

	"github.com/SafeAuth/backend/configs"
	"github.com/SafeAuth/backend/router"
)

var UnauthorizedErr = "you are not authorized to access this resource"

func main() {
	app := configs.FiberApp
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.Connect()

	router.ConnectRouter(app)

	app.Listen(":4000")
}
