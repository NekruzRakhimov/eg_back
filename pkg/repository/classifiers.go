package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetAllClassifiers() (classifiers []models.Classifier, err error) {
	sqlQuery := "SELECT id, name, full_name, code, to_char(created_at, 'DD.MM.YYYY') as created_at FROM classifiers WHERE is_removed=false"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&classifiers).Error; err != nil {
		return nil, err
	}

	return classifiers, err
}
