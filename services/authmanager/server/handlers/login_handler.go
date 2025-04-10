package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

func Login(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		oAuthType := c.Locals(middleware.O_AUTH_TYPE).(string)
		err := service.Login(oAuthType)
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
