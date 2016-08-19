package service

import "github.com/mdreizin/smartling/model"

type FileService interface {
	Pull(*FilePullParams) ([]*model.File, error)

	Push(*FilePushParams) (*model.FileStats, error)
}
