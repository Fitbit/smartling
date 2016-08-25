package service

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultAuthService_Authenticate(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	authToken := model.AuthToken{
		AccessToken:  "accessToken",
		RefreshToken: "refreshToken",
	}

	authService := DefaultAuthService{}

	resp, err := authService.Authenticate(&model.UserToken{
		ID:     "userId",
		Secret: "userSecret",
	})

	a := assert.New(t)

	a.NoError(err)
	a.EqualValues(&authToken, resp)
}

func TestDefaultAuthService_Refresh(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	authToken := model.AuthToken{
		AccessToken:  "accessToken",
		RefreshToken: "refreshToken",
	}

	authService := DefaultAuthService{}

	resp, err := authService.Refresh("refreshToken")

	a := assert.New(t)

	a.NoError(err)
	a.EqualValues(&authToken, resp)
}
