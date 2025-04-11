package services

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/oauth"
)

type LoginService interface {
	Login(oAuthType string) (string, error)
	OAuthCallback(oAuthType string, code string) ([]byte, error)
}

type loginService struct {
	log        *logger.Logger
	repository dbmanager.DbManagerClient
}

func NewLoginService(log *logger.Logger, repository dbmanager.DbManagerClient) LoginService {
	return &loginService{
		log:        log,
		repository: repository,
	}
}

func (ls *loginService) Login(oAuthType string) (string, error) {
	ls.log.Info("Login with ", oAuthType)

	switch oAuthType {
	case middleware.GOOGLE_OAUTH:
		googleConfig := oauth.GoogleConfigInit()
		url := googleConfig.AuthCodeURL(middleware.GOOGLE_OAUTH)
		return url, nil
	default:
		return "", fmt.Errorf("Unimplemented oAuthType: %s", oAuthType)
	}
}

func (ls *loginService) OAuthCallback(oAuthType string, code string) ([]byte, error) {
	ls.log.Info("OAuthCallback with ", oAuthType)
	switch oAuthType {
	case middleware.GOOGLE_OAUTH:
		// ver uma maneira de colocar esse bloco ed código em outra função através de interfaces
		googleConfig := oauth.GoogleConfigInit()
		token, err := googleConfig.Exchange(context.Background(), code)
		if err != nil {
			return nil, fmt.Errorf("Failed to exchange code for token: %v", err)
		}
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("Failed to get userinfo: %v", err)
		}

		userData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("JSON Parsing Failed: %v", err)
		}

		return userData, nil
	default:
		return nil, fmt.Errorf("Unimplemented oAuthType: %s", oAuthType)
	}
}
