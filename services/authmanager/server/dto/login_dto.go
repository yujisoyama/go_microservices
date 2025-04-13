package authmanagerdto

type OAuthLoginOutputDto struct {
	AccessToken string `json:"access_token"`
}

type MeOutputDto struct {
	UserId  string `json:"user_id"`
	OAuthId string `json:"oauth_id"`
}
