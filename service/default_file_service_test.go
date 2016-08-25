package service

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFileService_Push(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	fileService := DefaultFileService{}

	stats, err := fileService.Push(&FilePushParams{
		ProjectID:  "projectId",
		FilePath:   "../test/testdata/foo/en-US.json",
		Directives: map[string]string{},
		AuthToken:  "authToken",
	})

	a := assert.New(t)

	a.NoError(err)
	a.EqualValues(&model.FileStats{
		OverWritten: true,
		StringCount: 10,
		WordCount:   10,
	}, stats)
}

func TestDefaultFileService_Pull(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	fileService := DefaultFileService{}

	files, err := fileService.Pull(&FilePullParams{
		ProjectID: "projectId",
		FileURIs:  []string{},
		LocaleIDs: []string{"de-DE"},
	})

	a := assert.New(t)

	a.NoError(err)
	a.NotNil(files)
	a.Len(files, 1)
}
