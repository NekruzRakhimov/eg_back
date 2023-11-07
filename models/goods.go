package models

type Good struct {
	ID          int    `json:"id"`
	ExtID       string `json:"ext_id"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Classifier  string `json:"classifier"`
}
