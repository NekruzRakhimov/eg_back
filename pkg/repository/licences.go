package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetAllLicencesByEnterpriseID(enterpriseID int) (licences []models.Licence, err error) {
	if err = db.GetDBConn().
		Raw("SELECT * FROM licences WHERE enterprise_id = ?", enterpriseID).
		Scan(&licences).Error; err != nil {
		return nil, err
	}

	return licences, nil
}

func GetLicenceByID(id int) (l models.Licence, err error) {
	if err = db.GetDBConn().
		Raw("SELECT * FROM licences WHERE id = ?", id).
		Scan(&l).Error; err != nil {
		return models.Licence{}, err
	}

	return l, nil
}

func AddLicenceToEnterprise(enterpriseID int, licence models.Licence) error {
	sqlQuery := `INSERT INTO licences (enterprise_id, 
									  series, 
									  valid_until, 
									  issued_to, 
									  activity_type, 
									  licensee_legal_address,
									  issued_at, 
									  commission_protocol_number,
									  register_number,
                      				  iqtibos_info)
                 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	if err := db.GetDBConn().Exec(sqlQuery,
		enterpriseID,
		licence.Series,
		licence.ValidUntil,
		licence.IssuedTo,
		licence.ActivityType,
		licence.LicenseeLegalAddress,
		licence.IssuedAt,
		licence.CommissionProtocolNumber,
		licence.RegisterNumber,
		licence.IqtibosInfo).
		Omit("created_at, updated_at, deleted_at").Error; err != nil {
		return err
	}

	return nil
}
