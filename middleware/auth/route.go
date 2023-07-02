package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/")

	group.Post("/register", Register)
	group.Post("/login", Login)
	group.Get("/status", Status)

}
