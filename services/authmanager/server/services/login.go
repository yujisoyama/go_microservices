package services

import (
	// "context"
	// "fmt"
	// "io"
	// "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/oauth"
)

type LoginService interface {
	Login(oAuthType middleware.OAuthType) (string, error)
	// OAuthCallback(oAuthType string, code string) ([]byte, error)
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

func (ls *loginService) Login(oAuthType middleware.OAuthType) (string, error) {
	ls.log.Info("Login with ", oAuthType)
	oAuthConfig, exists := ls.oAuthConfigs[oAuthType]
	if !exists {
		return "", fiber.ErrNotImplemented
	}

	url := oAuthConfig.OAuthLogin()
	return url, nil
}

// func (ls *loginService) OAuthCallback(oAuthType string, code string) ([]byte, error) {
// 	ls.log.Info("OAuthCallback with ", oAuthType)
// 	switch oAuthType {
// 	case middleware.GOOGLE_OAUTH:
// 		// ver uma maneira de colocar esse bloco ed código em outra função através de interfaces
// 		googleConfig := oauth.GoogleConfigInit()
// 		token, err := googleConfig.Exchange(context.Background(), code)
// 		if err != nil {
// 			return nil, fmt.Errorf("Failed to exchange code for token: %v", err)
// 		}
// 		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
// 		if err != nil {
// 			return nil, fmt.Errorf("Failed to get userinfo: %v", err)
// 		}

// 		userData, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, fmt.Errorf("JSON Parsing Failed: %v", err)
// 		}

// 		return userData, nil
// 	default:
// 		return nil, fmt.Errorf("Unimplemented oAuthType: %s", oAuthType)
// 	}
// }
