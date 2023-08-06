package models

type Report struct {
	Active         int `json:"active"`
	Liquidated     int `json:"liquidated"`
	Stopped        int `json:"stopped"`
	OnBankruptcy   int `json:"on_bankruptcy"`
	LeasedOut      int `json:"leased_out"`
	ReOrganized    int `json:"re_organized"`
	CustomsDebtors int `json:"customs_debtors"`
	TaxDebtors     int `json:"tax_debtors"`
}

type ReportFilter struct {
	DateFrom          string `json:"date_from"`
	DateTo            string `json:"date_to"`
	Location          string `json:"location"`
	EconomicActivity  string `json:"economic_activity"`
	OrganizationAge   int    `json:"organization_age"`
	YearlyTurnover    int    `json:"yearly_turnover"`
	AuthorizedCapital int    `json:"authorized_capital"`
}
