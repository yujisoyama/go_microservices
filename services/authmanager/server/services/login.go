package services

import (
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
)

type LoginService interface {
	Login(oAuthType string) error
}

type loginService struct {
	log        *logger.Logger
	repository *dbmanager.DbManagerClient
}

func NewLoginService(log *logger.Logger) LoginService {
	return &loginService{
		log:        log,
		// repository: repository,
	}
}

func (ls *loginService) Login(oAuthType string) error {
	ls.log.Info("Login with ", oAuthType)
	return nil
}
