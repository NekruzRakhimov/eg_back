package repository

import (
	"eg_back/db"
	"eg_back/models"
)

func GetAllEmployeesByEnterpriseID(enterpriseID int) (employees []models.Employee, err error) {
	if err = db.GetDBConn().
		Raw("SELECT * FROM employees WHERE enterprise_id = ?", enterpriseID).
		Scan(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func GetEmployeeByID(id int) (e models.Employee, err error) {
	if err = db.GetDBConn().
		Raw("SELECT * FROM employees WHERE id = ?", id).
		Scan(&e).Error; err != nil {
		return models.Employee{}, err
	}

	return e, nil
}
