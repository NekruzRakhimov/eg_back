package models

type Enterprise struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	OwnershipType     string  `json:"ownership_type"`
	Okved             string  `json:"okved"`
	INN               string  `json:"inn"`
	EIN               string  `json:"ein"`
	RegistrationDate  string  `json:"registration_date"`
	Address           string  `json:"address"`
	AuthorizedCapital float64 `json:"authorized_capital"`
	OkvedID           int     `json:"okved_id"`
	KfcID             int     `json:"kfc_id"`
	//Rating      struct {
	//	Rate  float64 `json:"rate"`
	//	Count int     `json:"count"`
	//} `json:"rating"`
}
