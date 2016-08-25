package service

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/rest"
	"gopkg.in/resty.v0"
)

type DefaultAuthService struct{}

func (s *DefaultAuthService) Authenticate(userToken *model.UserToken) (*model.AuthToken, error) {
	return s.getToken(rest.StaticURL(rest.AuthenticateURL), userToken)
}

func (s *DefaultAuthService) Refresh(refreshToken string) (*model.AuthToken, error) {
	req := struct {
		RefreshToken string `json:"refreshToken"`
	}{
		refreshToken,
	}

	return s.getToken(rest.StaticURL(rest.AuthenticateRefreshURL), req)
}

func (s *DefaultAuthService) getToken(url string, req interface{}) (*model.AuthToken, error) {
	var (
		err  error
		resp *resty.Response
	)

	authToken := &model.AuthToken{}
	client := resty.R().SetBody(req).SetResult(rest.Model{}).SetError(rest.Model{})

	if resp, err = client.Post(url); err == nil {
		err = rest.Result(resp, &authToken)
	}

	return authToken, err
}
