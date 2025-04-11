package oauth

import (
	"fmt"

	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/authmanager/internal/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthConfig struct {
	OAuthConfig oauth2.Config
}

func GoogleConfigInit() OAuthInterface {
	googleConfig := oauth2.Config{
		ClientID:     utils.GetEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: utils.GetEnv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  fmt.Sprintf("%s/oauth-callback", utils.GetEnv("AUTH_MANAGER_HOST")),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return &GoogleOAuthConfig{
		OAuthConfig: googleConfig,
	}
}

func (g *GoogleOAuthConfig) OAuthLogin() string {
	url := g.OAuthConfig.AuthCodeURL(string(middleware.GOOGLE_OAUTH))
	return url
}
