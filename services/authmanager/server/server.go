package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/routes"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

type AuthManager struct {
	log     *logger.Logger
	Configs *AuthManagerConfigs
	app     *fiber.App
}

func NewAuthManager() *AuthManager {
	authManager := &AuthManager{
		log:     logger.NewLogger(),
		Configs: &AuthManagerConfigs{},
	}

	authManager.SetConfigs()
	return authManager
}

func (am *AuthManager) Run(ctx context.Context) error {
	am.app = fiber.New(fiber.Config{
		AppName: "AuthManager",
	})
	authMiddleware := middleware.NewAuthMiddleware()
	loginService := services.NewLoginService(am.log)

	am.app.Use(authMiddleware.CheckAuth())
	routes.LoginRouter(am.app, loginService)

	return am.app.Listen(fmt.Sprintf(":%s", am.Configs.port))
}
