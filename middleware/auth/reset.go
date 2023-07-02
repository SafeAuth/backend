package auth

import (
	"crypto/tls"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/SafeAuth/backend/database"
	"github.com/SafeAuth/backend/handler"

	mail "github.com/xhit/go-simple-mail/v2"
)

type ResetReq struct {
	Email string `json:"email" xml:"email"`
}

func Reset(c *fiber.Ctx) error {
	parser := new(ResetReq)
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	if parser.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Missing fields"})
	}

	token, err := database.Reset(parser.Email)
	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	server := mail.NewSMTPClient()

	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("EMAIL")
	server.Password = os.Getenv("EMAIL_PASSWORD")
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	smtpClient, err := server.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "message": "Failed to connect to SMTP server"})
	}

	email := mail.NewMSG()
	email.SetFrom("SafeAuth <no-reply@safeauth.me>").
		AddTo(parser.Email).
		SetSubject("SafeAuth Password Reset").
		SetBody("text/html", "<h1>SafeAuth Password Reset</h1><p>Click <a href=\"https://safeauth.me/reset/"+token+"\">here</a> to reset your password.</p>")
	err = email.Send(smtpClient)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "message": "Failed to send email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Password reset email sent"})

}
