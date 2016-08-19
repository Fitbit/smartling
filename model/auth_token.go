package model

type AuthToken struct {
	AccessToken  string `json:"accessToken" yaml:"AccessToken,omitempty"`
	RefreshToken string `json:"refreshToken" yaml:"-"`
}
