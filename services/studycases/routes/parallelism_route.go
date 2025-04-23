package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/services/studycases/server/handlers"
	"github.com/yujisoyama/go_microservices/services/studycases/server/services"
)

func ThreadRouter(app fiber.Router, service services.ThreadsService) {
	app.Get("/thread-tests/:testCase", handlers.Test(service))
}
