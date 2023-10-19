package models

type EnterpriseStructure struct {
	ID           int    `json:"id"`
	EnterpriseID int    `json:"enterprise_id"`
	ParentID     int    `json:"parent_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
}
