package rest

import (
	"gopkg.in/resty.v0"
)

func Result(resp *resty.Response, data interface{}) (err error) {
	var d *Model

	if resp.Result() != nil {
		d = resp.Result().(*Model)

		err = d.Data(data)
	}

	if err == nil {
		if (d != nil && !d.IsOK()) && resp.Error() != nil {
			err = resp.Error().(*Model).Error()
		}
	}

	return err
}
