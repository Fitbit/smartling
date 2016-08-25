package service

import "github.com/mdreizin/smartling/model"

type AuthService interface {
	Authenticate(userToken *model.UserToken) (*model.AuthToken, error)

	Refresh(refreshToken string) (*model.AuthToken, error)
}
