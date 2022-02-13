package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokensResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
