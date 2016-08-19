package rest

var (
	AuthenticateURL        = "/auth-api/v2/authenticate"
	AuthenticateRefreshURL = "/auth-api/v2/authenticate/refresh"
	FilePushURL            = "/files-api/v2/projects/{{ .ProjectID }}/file"
	FilePullURL            = "/files-api/v2/projects/{{ .ProjectID }}/files/zip"
)
