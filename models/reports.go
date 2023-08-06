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
