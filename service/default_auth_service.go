// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package service

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/rest"
	"gopkg.in/resty.v0"
)

type DefaultAuthService struct {
	Client *resty.Client `inject:"DefaultRestClient"`
}

func (s *DefaultAuthService) Authenticate(userToken *model.UserToken) (*model.AuthToken, error) {
	return s.getToken(rest.AuthenticateURL, userToken)
}

func (s *DefaultAuthService) Refresh(refreshToken string) (*model.AuthToken, error) {
	data := struct {
		RefreshToken string `json:"refreshToken"`
	}{
		refreshToken,
	}

	return s.getToken(rest.AuthenticateRefreshURL, data)
}

func (s *DefaultAuthService) getToken(url string, data interface{}) (*model.AuthToken, error) {
	var (
		err  error
		resp *resty.Response
	)

	authToken := &model.AuthToken{}
	req := s.Client.R().SetBody(data).SetResult(rest.Model{}).SetError(rest.Model{})

	if resp, err = req.Post(url); err == nil {
		err = rest.Result(resp, &authToken)
	}

	return authToken, err
}
