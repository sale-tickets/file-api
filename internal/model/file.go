package model

type File struct {
	Base
	Model    string `json:"model"`
	TargetId string `json:"target_id"`
	Path     string `json:"path"`
	FileType string `json:"file_type"`
}
