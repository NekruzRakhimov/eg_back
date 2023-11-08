package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetYearsOfEconomicActivity(enterpriseID int) (years []string, err error) {
	sqlQuery := `SELECT eea.year
				FROM enterprise_economic_activity eea
						 INNER JOIN public.enterprises e on e.id = eea.enterprise_id
				WHERE enterprise_id = ?
				GROUP BY eea.year
				ORDER BY eea.year`
	if err = db.GetDBConn().Raw(sqlQuery, enterpriseID).Pluck("year", &years).Error; err != nil {
		return nil, err
	}

	return years, nil
}

func GetEnterpriseActivityByYear(enterpriseID int, year string) (activity []models.EconomicActivity, err error) {
	sqlQuery := `SELECT eea.id, eea.enterprise_id, eea.key, value, eea.year, eea.created_at, eea.updated_at, eea.deleted_at, eea.is_removed
					FROM enterprise_economic_activity eea
							 INNER JOIN public.enterprises e on e.id = eea.enterprise_id
					WHERE year = ? AND enterprise_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, year, enterpriseID).Scan(&activity).Error; err != nil {
		return nil, err
	}

	return activity, nil
}
