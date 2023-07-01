package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/SafeAuth/backend/database"

	"github.com/SafeAuth/backend/configs"
	"github.com/SafeAuth/backend/router"
)

var UnauthorizedErr = "you are not authorized to access this resource"

// UnauthorizedErr is the error message that is returned when a user tries to
// access a resource that they are not authorized to access.

func main() {

	// Create a new Fiber app
	app := configs.FiberApp

	// Enable CORS
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	// This code loads the .env file and throws an error if it doesn't exist.
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Connect to the database
	database.Connect()

	// This code connects the router to the app.
	router.ConnectRouter(app)

	// This code starts the app on port 4000.
	app.Listen(":4000")
}
