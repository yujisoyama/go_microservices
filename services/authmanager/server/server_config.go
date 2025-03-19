package server

import "github.com/yujisoyama/go_microservices/pkg/utils"

type OAuthApiKeys string

const (
	GOOGLE_API_KEY   OAuthApiKeys = "GOOGLE_API_KEY"
	FACEBOOK_API_KEY OAuthApiKeys = "FACEBOOK_API_KEY"
	GITHUB_API_KEY   OAuthApiKeys = "GITHUB_API_KEY"
)

type OAuthConfigs struct {
	OAuthName    string
	ClientId     string
	ClientSecret string
}

type AuthManagerConfigs struct {
	port         string
	dbmApikey    string
	oAuthConfigs map[OAuthApiKeys]*OAuthConfigs
}

func (am *AuthManager) SetConfigs() {
	am.configs.port = utils.GetEnv("PORT")
	am.configs.dbmApikey = utils.GetEnv("DBM_API_KEY")
	am.configs.oAuthConfigs = map[OAuthApiKeys]*OAuthConfigs{
		GOOGLE_API_KEY: &OAuthConfigs{OAuthName: "Google", ClientId: utils.GetEnv("GOOGLE_CLIENT_ID"), ClientSecret: utils.GetEnv("GOOGLE_CLIENT_SECRET")},
	}
}
