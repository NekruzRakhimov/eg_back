package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetAllEmployeesByEnterpriseID(enterpriseID int) (employees []models.Employee, err error) {
	if err = db.GetDBConn().
		Raw("SELECT id, full_name, inn, job_title, birth_date, work_experience FROM employees WHERE enterprise_id = ?", enterpriseID).
		Scan(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func GetEmployeeByID(id int) (e models.Employee, err error) {
	if err = db.GetDBConn().
		Raw("SELECT id, full_name, inn, job_title, birth_date, work_experience FROM employees WHERE id = ?", id).
		Scan(&e).Error; err != nil {
		return models.Employee{}, err
	}

	return e, nil
}
