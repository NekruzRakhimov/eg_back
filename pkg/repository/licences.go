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
