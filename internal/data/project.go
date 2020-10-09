package data

import "github.com/google/uuid"

type Project struct {
	ProjectID   uuid.UUID `json:"projectId"`
	Commit      string    `json:"commit"`
	Name        string    `json:"name"`
	UnzipedPath string    `json:"unzip"`
	ZippedPath  string    `json:"zip"`
}
