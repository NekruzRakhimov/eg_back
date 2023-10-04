package repository

import (
	"eg_back/db"
	"eg_back/models"
	"fmt"
)

func GetAllClassifiers(filter models.Filter) (classifiers []models.Classifier, err error) {
	sqlQuery := "SELECT id, name, full_name, code, to_char(created_at, 'DD.MM.YYYY') as created_at FROM classifiers WHERE is_removed=false"
	if filter.Query != "" {
		sqlQuery = sqlQuery + " AND name iLIKE '%" + filter.Query + "%'"
	}

	if filter.Sort != "" {
		switch filter.Sort {
		case "name":
			sqlQuery = fmt.Sprintf("%s ORDER BY name", sqlQuery)
		case "id":
			sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
		default:
			sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
		}
	}

	if filter.Limit > 0 && filter.Page != 0 {
		sqlQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlQuery, filter.Limit, filter.Page-1)
	}
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&classifiers).Error; err != nil {
		return nil, err
	}

	return classifiers, err
}

func GetClassifierByCode(code string) (c models.Classifier, err error) {
	sqlQuery := `SELECT id, name, full_name, code, to_char(created_at, 'DD.MM.YYYY') as created_at
				 FROM classifiers 
				 WHERE is_removed=false AND lower(code) = lower(?)`
	if err = db.GetDBConn().Raw(sqlQuery, code).Scan(&c).Error; err != nil {
		return models.Classifier{}, err
	}

	return c, nil
}

func GetAllClassifiersCount(filter models.Filter) (count int, err error) {
	sqlQuery := "SELECT count(*) FROM classifiers WHERE is_removed=false "
	if filter.Query != "" {
		sqlQuery = sqlQuery + " AND name iLIKE '%" + filter.Query + "%'"
	}

	//if filter.Sort != "" {
	//	switch filter.Sort {
	//	case "name":
	//		sqlQuery = fmt.Sprintf("%s ORDER BY name", sqlQuery)
	//	case "id":
	//		sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
	//	default:
	//		sqlQuery = fmt.Sprintf("%s ORDER BY id", sqlQuery)
	//	}
	//}

	//if filter.Limit > 0 && filter.Page != 0 {
	//	sqlQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlQuery, filter.Limit, filter.Page-1)
	//}

	if err = db.GetDBConn().Raw(sqlQuery).Pluck("count", &count).Error; err != nil {
		return 0, err
	}

	return count, err
}

func GetClassifierItems(classifierID, ItemID int) (items []models.ClassifierItem, err error) {
	sqlQuery := `SELECT ci.id, c.name, ci.full_name, ci.code, to_char(ci.created_at, 'DD.MM.YYYY') as created_at
					FROM classifier_items ci
							 JOIN classifiers c on c.id = ci.classifier_id
					WHERE ci.parent_id = $1 AND ci.classifier_id = $2`
	if err = db.GetDBConn().Raw(sqlQuery, ItemID, classifierID).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
