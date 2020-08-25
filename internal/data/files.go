package data

type Language struct {
	Name  string   `json:"name,omitempty"`
	Files []string `json:"files,omitempty"`
}
