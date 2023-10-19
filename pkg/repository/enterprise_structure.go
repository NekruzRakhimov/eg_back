package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetEnterpriseStructure(enterpriseID int, structureID int) (s []models.EnterpriseStructure, err error) {
	sqlQuery := `SELECT id,
						   enterprise_id,
						   parent_id,
						   title,
						   description,
						   to_char(created_at, 'DD.MM.YYYY') as created_at
					FROM enterprise_structure
					WHERE enterprise_id = ? AND parent_id = ?;`
	if err = db.GetDBConn().Raw(sqlQuery, enterpriseID, structureID).Scan(&s).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func GetEnterpriseStructureByID(enterpriseID int, structureID int) (s models.EnterpriseStructure, err error) {
	sqlQuery := `SELECT id,
						   enterprise_id,
						   parent_id,
						   title,
						   description,
						   to_char(created_at, 'DD.MM.YYYY') as created_at
					FROM enterprise_structure
					WHERE enterprise_id = ? AND id = ?;`

	if err = db.GetDBConn().Raw(sqlQuery, enterpriseID, structureID).Scan(&s).Error; err != nil {
		return models.EnterpriseStructure{}, err
	}

	return s, nil
}
