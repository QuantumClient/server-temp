package models

type Project struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	GitHub  string `json:"github"`
	Link    NullString `json:"link"`
}
