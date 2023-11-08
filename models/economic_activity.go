package models

type BulkEconomicActivity struct {
	Year             string             `json:"year"`
	EconomicActivity []EconomicActivity `json:"economic_activity"`
}

type EconomicActivity struct {
	ID           int    `json:"id"`
	Key          string `json:"key"`
	Value        string `json:"value"`
	EnterpriseID int    `json:"enterprise_id"`
	Year         string `json:"year"`
}
