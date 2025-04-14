package authmanagerdto

type OAuthLoginOutputDto struct {
	AccessToken string `json:"access_token"`
}

type MeOutputDto struct {
	AccessToken   string `json:"access_token"`
	Id            string `json:"id"`
	OauthId       string `json:"oauth_id"`
	OauthtType    string `json:"oauth_type"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Picture       string `json:"picture"`
}
