package programs

import (
	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

type CreateReq struct {
	ProgramName string `json:"program_name" xml:"program_name"`
}

func Create(c *fiber.Ctx) error {
	parser := new(CreateReq)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	token := c.Cookies("token")
	verUser := VerUser(token)
	if !verUser.ValidUser {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "Unauthorized"})
	}

	if parser.ProgramName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Missing fields"})
	}

	encKey := uuid.New().String()
	err := database.CreateProgram(parser.ProgramName, encKey, verUser.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "message": "Failed to create program"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Program created"})
}
