package server

import "github.com/gofiber/fiber/v2"

func (am *AuthManager) Login(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}