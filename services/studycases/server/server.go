package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/services/studycases/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/studycases/routes"
	"github.com/yujisoyama/go_microservices/services/studycases/server/services"
)

type StudyCases struct {
	log     *logger.Logger
	configs *StudyCasesConfigs
	app     *fiber.App
}

func NewStudyCases() *StudyCases {
	authManager := &StudyCases{
		log:     logger.NewLogger(),
		configs: &StudyCasesConfigs{},
	}

	authManager.SetConfigs()
	return authManager
}

func (sc *StudyCases) Run(ctx context.Context) error {
	sc.app = fiber.New(fiber.Config{
		AppName: "StudyCases",
	})
	authMiddleware := middleware.NewAuthMiddleware()
	sc.app.Use(authMiddleware.CheckAuth())

	threadsService := services.NewThreadsService(sc.log)
	routes.ThreadRouter(sc.app, threadsService)

	return sc.app.Listen(fmt.Sprintf(":%s", sc.configs.port))
}
