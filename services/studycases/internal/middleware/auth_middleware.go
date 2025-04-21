package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/studycases/server/dto"
)

const (
	ME = "me"
)

type AuthMiddleware struct {
	log             *logger.Logger
	authManagerHost string
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		log:             logger.NewLogger(),
		authManagerHost: utils.GetEnv("AUTHMANAGER_HOST"),
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

		me, err := checkAuthManagerToken(c.Context(), am.authManagerHost, token)
		if err != nil {
			return utils.RestException(c, fiber.StatusUnauthorized, err.Error(), nil)
		}

		c.Locals(ME, me)
		return c.Next()
	}
}

func checkAuthManagerToken(ctx context.Context, authManagerHost string, token string) (*dto.MeDto, error) {
	url := fmt.Sprintf("%s/me", authManagerHost)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headers.ContentType, "application/json")
	req.Header.Set(headers.Authorization, fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error in me request with status code %d", resp.StatusCode)
	}

	respData, _ := io.ReadAll(resp.Body)
	var me dto.MeDto
	if err := json.Unmarshal(respData, &me); err != nil {
		return nil, err
	}

	return &me, nil
}
