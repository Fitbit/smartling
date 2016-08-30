package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
)

type pullResponse struct {
	Files   []*model.File
	Request *pullRequest
	Params  *service.FilePullParams
}
