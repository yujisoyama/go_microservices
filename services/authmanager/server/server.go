package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/grpc"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/routes"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

type AuthManager struct {
	log     *logger.Logger
	configs *AuthManagerConfigs
	app     *fiber.App
}

func NewAuthManager() *AuthManager {
	authManager := &AuthManager{
		log:     logger.NewLogger(),
		configs: &AuthManagerConfigs{},
	}

	authManager.SetConfigs()
	return authManager
}

func (am *AuthManager) Run(ctx context.Context) error {
	am.app = fiber.New(fiber.Config{
		AppName: "AuthManager",
	})
	authMiddleware := middleware.NewAuthMiddleware()

	grpcClient, err := grpc.InitGrpcClient(am.log, am.configs.dbmHost, am.configs.dbmApikey)
	if err != nil {
		am.log.Error("Failed to connect to DBManager", err)
		return err
	}

	loginService := services.NewLoginService(am.log, grpcClient)

	am.app.Use(authMiddleware.CheckAuth())
	routes.LoginRouter(am.app, loginService)

	return am.app.Listen(fmt.Sprintf(":%s", am.configs.port))
}
