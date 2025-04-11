package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

func Login(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		oAuthType := c.Locals(middleware.O_AUTH_TYPE).(middleware.OAuthType)
		url, err := service.Login(oAuthType)
		if err != nil {
			return utils.RestException(c, fiber.StatusNotImplemented, err.Error(), nil)
		}
		c.Status(fiber.StatusSeeOther)
		c.Redirect(url)
		return c.JSON(url)
	}
}

// func OAuthCallback(service services.LoginService) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		state := c.Query("state")
// 		code := c.Query("code")
// 		userData, err := service.OAuthCallback(state, code)
// 		if err != nil {
// 			return utils.RestException(c, fiber.StatusInternalServerError, err.Error(), nil)
// 		}

// 		return c.SendString(string(userData))
// 	}
// }
