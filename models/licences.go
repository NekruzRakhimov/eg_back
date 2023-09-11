package models

type Licence struct {
	ID                       int    `json:"id"`
	EnterpriseID             int    `json:"enterprise_id"`
	Series                   string `json:"series"`
	ValidUntil               string `json:"valid_until"`
	IssuedTo                 string `json:"issued_to"`
	ActivityType             string `json:"activity_type"`
	LicenseeLegalAddress     string `json:"licensee_legal_address"`
	IqtibosInfo              string `json:"iqtibos_info"`
	IssuedAt                 string `json:"issued_at"`
	CommissionProtocolNumber string `json:"commission_protocol_number"`
	RegisterNumber           string `json:"register_number"`
}
