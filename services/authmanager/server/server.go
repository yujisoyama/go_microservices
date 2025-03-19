package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
)

type AuthManager struct {
	log     *logger.Logger
	configs *AuthManagerConfigs
	app     *fiber.App
}

func NewAuthManager() *AuthManager {
	return &AuthManager{
		log:     logger.NewLogger(),
		configs: &AuthManagerConfigs{},
	}
}

func (am *AuthManager) Run(ctx context.Context) error {
	am.SetConfigs()

	am.app = fiber.New(fiber.Config{
		AppName: "AuthManager",
	})

	am.app.Post("/login", am.CheckApiKey, am.Login)

	return am.app.Listen(fmt.Sprintf(":%s", am.configs.port))
}
