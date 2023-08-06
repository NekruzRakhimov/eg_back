package repository

import "eg_back/db"

func GetReportByStatus(status string) (count int, err error) {
	sqlQuery := "SELECT count(id) as count FROM enterprises WHERE status = ?"
	if err = db.GetDBConn().Raw(sqlQuery, status).Pluck("count", &count).Error; err != nil {
		return 0, err
	}

	return count, err
}
