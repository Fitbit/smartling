package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
)

type pullRequest struct {
	Files                  []string
	Locales                []string
	IncludeOriginalStrings bool
	RetrievalType          string
	Config                 *model.ProjectConfig
	Resource               *model.ProjectResource
	AuthToken              string
	FileService            service.FileService
}
