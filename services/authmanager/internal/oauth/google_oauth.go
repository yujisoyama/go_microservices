package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yujisoyama/go_microservices/pkg/pb/user"
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

type GoogleUser struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func (googleUser *GoogleUser) ParseToEntity() *user.User {
	return &user.User{
		OauthId:       googleUser.Id,
		OauthType:     string(middleware.GOOGLE_OAUTH),
		Email:         googleUser.Email,
		VerifiedEmail: googleUser.VerifiedEmail,
		FirstName:     googleUser.GivenName,
		LastName:      googleUser.FamilyName,
		Picture:       googleUser.Picture,
	}
}

func (g *GoogleOAuthConfig) OAuthLogin() string {
	url := g.OAuthConfig.AuthCodeURL(string(middleware.GOOGLE_OAUTH))
	return url
}

func (g *GoogleOAuthConfig) OAuthCallback(code string) (*user.User, error) {
	token, err := g.OAuthConfig.Exchange(context.Background(), code)
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

	var googleUser GoogleUser
	if err := json.Unmarshal(userData, &googleUser); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal GoogleUser: %v", err)
	}

	user := googleUser.ParseToEntity()
	return user, nil
}
