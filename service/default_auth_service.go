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
	return s.getToken(rest.StaticURL(rest.AuthenticateRefreshURL), struct {
		RefreshToken string `json:"refreshToken"`
	}{
		refreshToken,
	})
}

func (s *DefaultAuthService) getToken(url string, data interface{}) (*model.AuthToken, error) {
	var (
		err  error
		resp *resty.Response
	)

	authToken := &model.AuthToken{}
	req := resty.R().SetBody(data).SetResult(rest.Model{}).SetError(rest.Model{})

	if resp, err = req.Post(url); err == nil {
		err = rest.Result(resp, &authToken)
	}

	return authToken, err
}
