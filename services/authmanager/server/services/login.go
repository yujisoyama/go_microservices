package services

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/oauth"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/jwt"
	authmanagerdto "github.com/yujisoyama/go_microservices/services/authmanager/server/dto"
)

type LoginService interface {
	Login(oAuthType middleware.OAuthType) (string, int, error)
	LoginCallback(oAuthType middleware.OAuthType, code string) (*authmanagerdto.OAuthLoginOutputDto, int, error)
}

type loginService struct {
	log          *logger.Logger
	repository   dbmanager.DbManagerClient
	oAuthConfigs map[middleware.OAuthType]oauth.OAuthInterface
	jwtService   *jwt.JWTService
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
		jwtService:   jwt.NewJWTConfigs(),
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

func (ls *loginService) LoginCallback(oAuthType middleware.OAuthType, code string) (*authmanagerdto.OAuthLoginOutputDto, int, error) {
	ls.log.Info("OAuthCallback with ", oAuthType)
	oAuthConfig, exists := ls.oAuthConfigs[oAuthType]
	if !exists {
		return nil, fiber.StatusNotImplemented, fmt.Errorf("oAuthType: %s not found", oAuthType)
	}

	user, err := oAuthConfig.OAuthCallback(code)
	if err != nil {
		return nil, fiber.StatusInternalServerError, fmt.Errorf("Error in OAuthCallback: %v", err)
	}

	upsertResp, err := ls.repository.UpsertUser(context.Background(), &dbmanager.UpsertUserRequest{
		User: user,
	})
	if err != nil {
		return nil, fiber.StatusInternalServerError, fmt.Errorf("Error in UpsertUser: %v", err)
	}

	token, err := ls.jwtService.GenerateToken(jwt.TokenInfo{
		UserId:  upsertResp.User.Id,
		OAuthId: upsertResp.User.OauthId,
	})
	if err != nil {
		return nil, fiber.StatusInternalServerError, fmt.Errorf("Error in GenerateToken: %v", err)
	}

	resp := &authmanagerdto.OAuthLoginOutputDto{
		AccessToken: token,
	}

	return resp, fiber.StatusOK, nil
}
