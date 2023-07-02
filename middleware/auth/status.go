package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SafeAuth/backend/database"
)

func Status(c *fiber.Ctx) error {
	token := c.Cookies("token")
	VerifyUser := database.VerUser(token)

	if !VerifyUser.ValidUser || VerifyUser.Banned {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "message": "not authorized", "auth": false})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":    false,
		"message":  "authorized",
		"auth":     true,
		"username": VerifyUser.Username,
		"admin":    VerifyUser.Admin,
		"uid":      VerifyUser.Uid,
	})

}
