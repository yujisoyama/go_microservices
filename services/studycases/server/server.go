package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
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

	paralellismService := services.NewParallelismService(sc.log)

	routes.ParallelismRouter(sc.app, paralellismService)

	return sc.app.Listen(fmt.Sprintf(":%s", sc.configs.port))
}
