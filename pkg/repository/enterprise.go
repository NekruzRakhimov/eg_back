package repository

import (
	"eg_back/db"
	"eg_back/models"
	"fmt"
)

func GetAllEnterprises(filter models.Filter) (e []models.Enterprise, err error) {
	fmt.Println("query", filter.Query)
	sqlQuery := "select * from public.enterprises WHERE true "
	if filter.Query != "" {
		sqlQuery = sqlQuery + " AND name iLIKE '%" + filter.Query + "%'"
	}

	if filter.AuthorizedCapitalFilter != -1 && filter.AuthorizedCapitalFilter != 0 {
		sqlQuery = fmt.Sprintf("%s AND authorized_capital < %d",
			sqlQuery, filter.AuthorizedCapitalFilter)
	}

	if filter.EnterpriseAgeFilter != "" && filter.EnterpriseAgeFilter != "-1" {
		sqlQuery = fmt.Sprintf("%s AND date_part('year', CURRENT_DATE) - date_part('year', created_at) %s", sqlQuery, filter.EnterpriseAgeFilter)
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
	fmt.Println("query", filter.Query)
	sqlQuery := "select * from public.enterprises WHERE true "
	if filter.Query != "" {
		sqlQuery = sqlQuery + " AND name iLIKE '%" + filter.Query + "%'"
	}

	if filter.AuthorizedCapitalFilter != -1 && filter.AuthorizedCapitalFilter != 0 {
		sqlQuery = fmt.Sprintf("%s AND authorized_capital < %d",
			sqlQuery, filter.AuthorizedCapitalFilter)
	}

	//if filter.Sort != "" {
	//	switch filter.Sort {
	//	case "name":
	//		sqlQuery = fmt.Sprintf("%s ORDER BY name", sqlQuery)
	//	case "id":
	//		sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
	//	case "authorized_capital":
	//		sqlQuery = fmt.Sprintf("%s ORDER BY authorized_capital", sqlQuery)
	//	default:
	//		sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
	//	}
	//}

	//if filter.Limit > 0 && filter.Page != 0 {
	//	sqlQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlQuery, filter.Limit, filter.Page-1)
	//}

	if err = db.GetDBConn().Raw("select count(id) from public.enterprises").Pluck("count", &count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetEnterpriseByID(id int) (e models.Enterprise, err error) {
	if err = db.GetDBConn().Raw("select * from public.enterprises WHERE id = ?", id).Scan(&e).Error; err != nil {
		return models.Enterprise{}, err
	}

	return e, nil
}
