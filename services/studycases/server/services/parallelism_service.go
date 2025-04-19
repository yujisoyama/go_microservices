package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
)

type ParallelismService interface {
	Test() (int, error)
}

type parallelismService struct {
	log *logger.Logger
}

func NewParallelismService(log *logger.Logger) ParallelismService {
	return &parallelismService{
		log: log,
	}
}

func (ps *parallelismService) Test() (int, error) {
	ps.log.Info("Parallelism Test")
	return fiber.StatusNoContent, nil
}
