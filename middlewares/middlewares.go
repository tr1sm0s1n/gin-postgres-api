package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Authority(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	expectedToken := "token"
	if authorizationHeader != "Bearer "+expectedToken {
		return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
	}

	return c.Next()
}
