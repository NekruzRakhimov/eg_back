package repository

import (
	"eg_back/db"
	"eg_back/models"
	"fmt"
)

func CreateEnterprise(e models.Enterprise) (err error) {
	if err = db.GetDBConn().Table("enterprises").Omit("id, created_at").Create(&e).Error; err != nil {
		return err
	}

	return nil
}

func GetAllEnterprises(filter models.Filter) (e []models.Enterprise, err error) {
	fmt.Println("query", filter.Query)
	sqlQuery := `select 
					id,
					   name,
					   (select full_name from classifier_items WHERE id = okved_id) as okved,
						(select full_name from classifier_items WHERE id = kfc_id) as ownership_type,
					   inn,
					   ein,
					   registration_date,
					   address,
					   authorized_capital,
					   is_removed,
					   status,
					   yearly_turnover
				from public.enterprises WHERE true `
	if filter.Query != "" {
		sqlQuery = sqlQuery + " AND name iLIKE '%" + filter.Query + "%'"
	}

	if filter.Location != "" {
		sqlQuery = sqlQuery + " AND address iLIKE '%" + filter.Location + "%'"
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
			sqlQuery = fmt.Sprintf("%s ORDER BY id DESC", sqlQuery)
		case "authorized_capital":
			sqlQuery = fmt.Sprintf("%s ORDER BY authorized_capital", sqlQuery)
		default:
			sqlQuery = fmt.Sprintf("%s ORDER BY id DESC", sqlQuery)
		}
	} else {
		sqlQuery = fmt.Sprintf("%s ORDER BY id DESC", sqlQuery)
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
	sqlQuery := `select 
					id,
									   name,
									   (select full_name from classifier_items WHERE id = okved_id) as okved,
										(select full_name from classifier_items WHERE id = kfc_id) as ownership_type,
									   inn,
									   ein,
									   registration_date,
									   address,
									   authorized_capital,
									   is_removed,
									   status,
									   yearly_turnover
				from public.enterprises WHERE true `
	if filter.Query != "" {
		sqlQuery = sqlQuery + " AND name iLIKE '%" + filter.Query + "%'"
	}

	if filter.Location != "" {
		sqlQuery = sqlQuery + " AND address iLIKE '%" + filter.Location + "%'"
	}

	if filter.AuthorizedCapitalFilter != -1 && filter.AuthorizedCapitalFilter != 0 {
		sqlQuery = fmt.Sprintf("%s AND authorized_capital < %d",
			sqlQuery, filter.AuthorizedCapitalFilter)
	}

	if filter.EnterpriseAgeFilter != "" && filter.EnterpriseAgeFilter != "-1" {
		sqlQuery = fmt.Sprintf("%s AND date_part('year', CURRENT_DATE) - date_part('year', created_at) %s", sqlQuery, filter.EnterpriseAgeFilter)
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
	if err = db.GetDBConn().Raw(`select id,
					   name,
					   (select full_name from classifier_items WHERE id = okved_id) as okved,
						(select full_name from classifier_items WHERE id = kfc_id) as ownership_type,
					   inn,
					   ein,
					   registration_date,
					   address,
					   authorized_capital,
					   is_removed,
					   status,
					   yearly_turnover from public.enterprises WHERE id = ?`, id).Scan(&e).Error; err != nil {
		return models.Enterprise{}, err
	}

	return e, nil
}
