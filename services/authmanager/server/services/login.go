package services

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/oauth"
)

type LoginService interface {
	Login(oAuthType middleware.OAuthType) (string, int, error)
	LoginCallback(oAuthType middleware.OAuthType, code string) ([]byte, int, error)
}

type loginService struct {
	log          *logger.Logger
	repository   dbmanager.DbManagerClient
	oAuthConfigs map[middleware.OAuthType]oauth.OAuthInterface
}

func NewLoginService(log *logger.Logger, repository dbmanager.DbManagerClient) LoginService {
	// start hashmap with oAuthConfigs
	oAuthConfigs := map[middleware.OAuthType]oauth.OAuthInterface{
		middleware.GOOGLE_OAUTH: oauth.GoogleConfigInit(),
	}

	return &loginService{
		log:          log,
		repository:   repository,
		oAuthConfigs: oAuthConfigs,
	}
}

func (ls *loginService) Login(oAuthType middleware.OAuthType) (string, int, error) {
	ls.log.Info("Login with ", oAuthType)
	oAuthConfig, exists := ls.oAuthConfigs[oAuthType]
	if !exists {
		return "", fiber.StatusNotImplemented, fmt.Errorf("oAuthType: %s not found", oAuthType)
	}

	url := oAuthConfig.OAuthLogin()
	return url, fiber.StatusOK, nil
}

func (ls *loginService) LoginCallback(oAuthType middleware.OAuthType, code string) ([]byte, int, error) {
	ls.log.Info("OAuthCallback with ", oAuthType)
	oAuthConfig, exists := ls.oAuthConfigs[oAuthType]
	if !exists {
		return nil, fiber.StatusNotImplemented, fmt.Errorf("oAuthType: %s not found", oAuthType)
	}

	user, err := oAuthConfig.OAuthCallback(code)
	if err != nil {
		return nil, fiber.StatusInternalServerError, fmt.Errorf("Error in OAuthCallback: %v", err)
	}

	_, err = ls.repository.UpsertUser(context.Background(), &dbmanager.UpsertUserRequest{
		User: user,
	})
	if err != nil {
		return nil, fiber.StatusInternalServerError, fmt.Errorf("Error in UpsertUser: %v", err)
	}

	return nil, fiber.StatusOK, nil
}
