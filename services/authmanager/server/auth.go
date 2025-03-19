package server

import (
	"fmt"

	"github.com/go-http-utils/headers"
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
)

func (am *AuthManager) CheckApiKey(c *fiber.Ctx) error {
	auth := c.Get(headers.Authorization)
	if len(auth) < 1 {
		return utils.RestException(c, fiber.StatusUnauthorized, "Missing api key", nil)
	}

	var token OAuthApiKeys
	if _, err := fmt.Sscanf(auth, "Bearer %s", &token); err != nil {
		utils.RestException(c, fiber.StatusUnauthorized, "Invalid Authorization format", nil)
	}

	if _, exists := am.configs.oAuthConfigs[token]; !exists {
		return utils.RestException(c, fiber.StatusUnauthorized, "Invalid api key", nil)
	}

	return c.Next()
}
