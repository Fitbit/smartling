package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
)

type pushRequest struct {
	Path        string
	Config      *model.ProjectConfig
	Resource    *model.ProjectResource
	AuthToken   string
	FileService service.FileService
}
