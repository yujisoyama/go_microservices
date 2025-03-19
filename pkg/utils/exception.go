package utils

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcException(code codes.Code, customMsg string) error {
	return status.Error(codes.Internal, customMsg)
}

func RestException(c *fiber.Ctx, code int, customMsg string, err error) error {
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.Status(code).JSON(fiber.Map{
		"error": customMsg,
	})
}
