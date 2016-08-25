package service

type FilePushParams struct {
	ProjectID  string            `url:"-"`
	FileURI    string            `url:"fileUri"`
	FilePath   string            `url:"-"`
	FileType   string            `url:"fileType"`
	Authorize  bool              `url:"authorize,omitempty"`
	Directives map[string]string `url:"-"`
	AuthToken  string            `url:"-"`
}
