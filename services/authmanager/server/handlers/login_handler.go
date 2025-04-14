package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/jwt"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

func OAuthLogin(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		oAuthType := c.Locals(middleware.O_AUTH_TYPE).(middleware.OAuthType)
		url, code, err := service.OAuthLogin(oAuthType)
		if err != nil {
			return utils.RestException(c, code, err.Error(), nil)
		}
		c.Status(fiber.StatusSeeOther)
		c.Redirect(url)
		return c.JSON(url)
	}
}

func OAuthCallback(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		state := c.Query("state")
		code := c.Query("code")
		userData, httpCode, err := service.OAuthLoginCallback(middleware.OAuthType(state), code)
		if err != nil {
			return utils.RestException(c, httpCode, err.Error(), nil)
		}

		return c.Status(httpCode).JSON(userData)
	}
}

func Me(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenInfo := c.Locals(middleware.TOKEN_INFO).(*jwt.TokenInfo)
		token := c.Locals(middleware.ACCESS_TOKEN).(string)
		user, code, err := service.Me(token, *tokenInfo)
		if err != nil {
			return utils.RestException(c, code, err.Error(), nil)
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}
