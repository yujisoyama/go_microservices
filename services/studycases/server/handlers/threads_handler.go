package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/studycases/server/services"
)

func Test(service services.ThreadsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		testCase := c.Params("testCase")
		t, code, err := service.TestThreads(testCase)
		if err != nil {
			return utils.RestException(c, code, err.Error(), nil)
		}
		return c.Status(code).SendString(t)
	}
}