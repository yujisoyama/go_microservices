package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/studycases/server/services"
)

func Test(service services.ParallelismService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code, err := service.Test()
		if err != nil {
			return utils.RestException(c, code, err.Error(), nil)
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}