package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SafeAuth/backend/database"
	"github.com/SafeAuth/backend/handler"
)

type RegisterReq struct {
	Email    string `json:"email" xml:"email"`
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func Register(c *fiber.Ctx) error {
	parser := new(RegisterReq)

	if err := c.BodyParser(parser); err != nil {
		return err
	}

	if parser.Username == "" || parser.Password == "" || parser.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Missing fields"})
	}

	token, err := database.Register(parser.Username, parser.Email, parser.Password, c.IP())
	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	c.Cookie(&fiber.Cookie{
		Domain:   ".safeauth.me",
		Path:     "/",
		Name:     "token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "user registered"})

}
