package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/handlers"
	"github.com/yujisoyama/go_microservices/services/authmanager/server/services"
)

func LoginRouter(app fiber.Router, service services.LoginService) {
	app.Get("/oauth-login", handlers.Login(service))
	app.Get("/oauth-callback", handlers.OAuthCallback(service))
}
