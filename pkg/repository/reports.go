package repository

import (
	"eg_back/db"
	"eg_back/models"
	"fmt"
)

func GetReportByStatus(status string, filter models.ReportFilter) (count int, err error) {
	sqlQuery := "SELECT count(id) as count FROM enterprises WHERE status = ? "

	if filter.DateFrom != "" {
		sqlQuery = fmt.Sprintf(`%s AND 
		to_date(to_char(created_at, 'DD-MM-YYYY'), 'DD-MM-YYYY') >= to_date('%s', 'YYYY-MM-DD')`, sqlQuery, filter.DateFrom)
	}

	if filter.DateTo != "" {
		sqlQuery = fmt.Sprintf(`%s AND 
		to_date(to_char(created_at, 'DD-MM-YYYY'), 'DD-MM-YYYY') <= to_date('%s', 'YYYY-MM-DD')`, sqlQuery, filter.DateTo)
	}

	if filter.Location != "" {
		sqlQuery += "AND address ilike '%" + filter.Location + "%'"
	}

	if filter.EconomicActivity != "" {
		sqlQuery += "AND okved ilike '%" + filter.EconomicActivity + "%'"
	}

	if filter.OrganizationAge > 0 {
		sqlQuery = fmt.Sprintf("%s AND date_part('year', CURRENT_DATE) - date_part('year', created_at) >= %d", sqlQuery, filter.OrganizationAge)
	}

	if filter.YearlyTurnover > 0 {
		sqlQuery = fmt.Sprintf("%s AND yearly_turnover >= %d", sqlQuery, filter.YearlyTurnover)
	}

	if filter.AuthorizedCapital > 0 {
		sqlQuery = fmt.Sprintf("%s AND authorized_capital >= %d", sqlQuery, filter.AuthorizedCapital)
	}

	if err = db.GetDBConn().Raw(sqlQuery, status).Pluck("count", &count).Error; err != nil {
		return 0, err
	}

	return count, err
}
