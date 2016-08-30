package rest

import "gopkg.in/resty.v0"

func Client() *resty.Client {
	return resty.New().SetHostURL(BaseURL)
}
