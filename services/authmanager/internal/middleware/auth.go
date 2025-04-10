package middleware

import (
	"fmt"

	"github.com/go-http-utils/headers"
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
)

const (
	O_AUTH_TYPE  = "oAuthType"
	GOOGLE_OAUTH = "GOOGLE_OAUTH"
	GITHUB_OAUTH = "GITHUB_OAUTH"
)

type AuthMiddleware struct {
	apiKeys map[string]string
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		apiKeys: map[string]string{
			utils.GetEnv("GOOGLE_API_KEY"): GOOGLE_OAUTH,
			utils.GetEnv("GITHUB_API_KEY"): GITHUB_OAUTH,
		},
	}
}

func (am *AuthMiddleware) CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(headers.Authorization)
		if len(auth) < 1 {
			return utils.RestException(c, fiber.StatusUnauthorized, "Missing api key", nil)
		}

		var token string
		if _, err := fmt.Sscanf(auth, "Bearer %s", &token); err != nil {
			utils.RestException(c, fiber.StatusUnauthorized, "Invalid Authorization format", nil)
		}

		oAuth, exists := am.apiKeys[token]
		if exists {
			c.Locals(O_AUTH_TYPE, oAuth)
			return c.Next()
		}

		return utils.RestException(c, fiber.StatusUnauthorized, "Invalid api key", nil)
	}
}
