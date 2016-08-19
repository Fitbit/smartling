package model

type FileStats struct {
	OverWritten bool   `json:"overWritten"`
	StringCount uint64 `json:"stringCount"`
	WordCount   uint64 `json:"wordCount"`
}
