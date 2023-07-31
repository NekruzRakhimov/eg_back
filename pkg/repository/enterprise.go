package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetAllEnterprises() (e []models.Enterprise, err error) {
	if err = db.GetDBConn().Raw("select * from public.enterprises").Scan(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}
