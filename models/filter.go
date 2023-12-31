package models

type Filter struct {
	Page                    int    `json:"page"`
	Limit                   int    `json:"limit"`
	Query                   string `json:"query"`
	Sort                    string `json:"sort"`
	AuthorizedCapitalFilter int    `json:"authorized_capital_filter"`
	EnterpriseAgeFilter     string `json:"enterprise_age_filter"`
	Location                string `json:"location"`
}
