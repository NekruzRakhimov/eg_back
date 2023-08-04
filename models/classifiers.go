package models

type Classifier struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
}

type ClassifierItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
}
