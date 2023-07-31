package repository

import (
	"eg_back/db"
	"eg_back/models"
	"fmt"
)

func GetAllEnterprises(filter models.Filter) (e []models.Enterprise, err error) {
	fmt.Println("query", filter.Query)
	sqlQuery := "select * from public.enterprises"
	if filter.Query != "" {
		sqlQuery = sqlQuery + " WHERE name LIKE '%" + filter.Query + "%'"
	}

	if filter.Sort != "" {
		switch filter.Sort {
		case "name":
			sqlQuery = fmt.Sprintf("%s ORDER BY name", sqlQuery)
		case "id":
			sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
		case "authorized_capital":
			sqlQuery = fmt.Sprintf("%s ORDER BY authorized_capital", sqlQuery)
		default:
			sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
		}
	}

	if filter.Limit > 0 && filter.Page != 0 {
		sqlQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlQuery, filter.Limit, filter.Page-1)
	}

	if err = db.GetDBConn().Raw(sqlQuery).Scan(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func GetEnterprisesCount(filter models.Filter) (count int, err error) {
	if err = db.GetDBConn().Raw("select count(id) from public.enterprises").Pluck("count", &count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
