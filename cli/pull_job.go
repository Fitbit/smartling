package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/go-playground/pool.v3"
)

func pullJob(req *pullRequest) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		uris := []string{}

		for _, path := range req.Files {
			logInfo(fmt.Sprintf("%s", color.MagentaString(path)))

			uris = append(uris, req.Config.FileURI(path))
		}

		params := &service.FilePullParams{
			ProjectID:              req.Config.Project.ID,
			FileURIs:               uris,
			LocaleIDs:              req.Locales,
			RetrievalType:          req.RetrievalType,
			IncludeOriginalStrings: req.IncludeOriginalStrings,
			AuthToken:              req.AuthToken,
		}

		files, err := req.FileService.Pull(params)

		return &pullResponse{
			Files:   files,
			Request: req,
		}, err
	}
}
