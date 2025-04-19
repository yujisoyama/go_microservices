package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/services/studycases/server/handlers"
	"github.com/yujisoyama/go_microservices/services/studycases/server/services"
)

func ParallelismRouter(app fiber.Router, service services.ParallelismService) {
	app.Get("/parallelism-test", handlers.Test(service))
}
