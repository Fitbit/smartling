package service

type FilePullParams struct {
	ProjectID              string   `url:"-"`
	FileURIs               []string `url:"fileUris,brackets"`
	LocaleIDs              []string `url:"localeIds,brackets"`
	RetrievalType          string   `url:"retrievalType,omitempty"`
	IncludeOriginalStrings bool     `url:"includeOriginalStrings,omitempty"`
	AuthToken              string   `url:"-"`
}
