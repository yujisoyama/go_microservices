package oauth

import "github.com/yujisoyama/go_microservices/pkg/pb/user"

type OAuthInterface interface {
	OAuthLogin() string
	OAuthCallback(code string) (*user.User, error)
}
