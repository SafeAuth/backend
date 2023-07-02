package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SafeAuth/backend/database"
	"github.com/SafeAuth/backend/handler"
)

type LoginReq struct {
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func Login(c *fiber.Ctx) error {
	parser := new(LoginReq)
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	token, err := database.Login(parser.Username, parser.Password, c.IP())
	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"error": false, "message": "logged in"})

}
