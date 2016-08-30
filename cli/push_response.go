package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
)

type pushResponse struct {
	Stats  *model.FileStats
	Params *service.FilePushParams
}
