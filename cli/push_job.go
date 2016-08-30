package main

import (
	"github.com/mdreizin/smartling/service"
	"gopkg.in/go-playground/pool.v3"
)

func pushJob(req *pushRequest) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		directives := req.Resource.Directives.WithPrefix()

		params := &service.FilePushParams{
			ProjectID:  req.Config.Project.ID,
			FileURI:    req.Config.FileURI(req.Path),
			FilePath:   req.Path,
			FileType:   req.Resource.Type,
			Authorize:  req.Resource.AuthorizeContent,
			Directives: directives,
			AuthToken:  req.AuthToken,
		}

		stats, err := req.FileService.Push(params)

		return &pushResponse{
			Stats:  stats,
			Params: params,
		}, err
	}
}
