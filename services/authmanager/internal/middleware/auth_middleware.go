package middleware

import (
	"fmt"

	"github.com/go-http-utils/headers"
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/jwt"
)

// create to define the oAuthTypes
type OAuthType string

const (
	O_AUTH_TYPE            = "oAuthType"
	ACCESS_TOKEN           = "accessToken"
	TOKEN_INFO             = "tokenInfo"
	GOOGLE_OAUTH OAuthType = "GOOGLE_OAUTH"
	GITHUB_OAUTH OAuthType = "GITHUB_OAUTH"
)

type AuthMiddleware struct {
	log        *logger.Logger
	apiKeys    map[string]OAuthType
	jwtService *jwt.JWTService
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		log: logger.NewLogger(),
		apiKeys: map[string]OAuthType{
			utils.GetEnv("AUTH_MANAGER_GOOGLE_API_KEY"): GOOGLE_OAUTH,
			utils.GetEnv("AUTH_MANAGER_GITHUB_API_KEY"): GITHUB_OAUTH,
		},
		jwtService: jwt.NewJWTConfigs(),
	}
}

func (am *AuthMiddleware) CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(headers.Authorization)
		if len(auth) < 1 {
			return utils.RestException(c, fiber.StatusUnauthorized, "Missing token", nil)
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

		tokenInfo, err := am.jwtService.VerifyToken(token)
		if err != nil {
			return utils.RestException(c, fiber.StatusUnauthorized, err.Error(), nil)
		}

		c.Locals(ACCESS_TOKEN, token)
		c.Locals(TOKEN_INFO, tokenInfo)
		return c.Next()
	}
}
