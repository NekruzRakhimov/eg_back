package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetAllGoodsByEnterpriseID(enterpriseID int) (goods []models.Good, err error) {
	if err = db.GetDBConn().
		Raw(`SELECT g.id, g.ext_id, g.full_name, g.description, ci.full_name as classifier
				FROM goods g
				INNER JOIN public.classifier_items ci on g.classifier_id = ci.id WHERE enterprise_id = ?`, enterpriseID).
		Scan(&goods).Error; err != nil {
		return nil, err
	}

	return goods, nil
}
