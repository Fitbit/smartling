package model

type UserToken struct {
	ID     string `json:"userIdentifier" yaml:"UserId,omitempty"`
	Secret string `json:"userSecret" yaml:"UserSecret,omitempty"`
}
