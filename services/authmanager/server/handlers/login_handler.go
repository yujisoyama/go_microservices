package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/jwt"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	authmanagerdto "github.com/yujisoyama/go_microservices/services/authmanager/server/dto"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

func Login(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		oAuthType := c.Locals(middleware.O_AUTH_TYPE).(middleware.OAuthType)
		url, code, err := service.Login(oAuthType)
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
		userData, httpCode, err := service.LoginCallback(middleware.OAuthType(state), code)
		if err != nil {
			return utils.RestException(c, httpCode, err.Error(), nil)
		}

		return c.Status(httpCode).JSON(userData)
	}
}

func Me(service services.LoginService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenInfo := c.Locals(middleware.TOKEN_INFO).(*jwt.TokenInfo)
		resp := authmanagerdto.MeOutputDto{
			UserId:  tokenInfo.UserId,
			OAuthId: tokenInfo.OAuthId,
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
