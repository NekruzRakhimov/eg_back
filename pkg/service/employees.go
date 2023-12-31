package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllEmployeesByEnterpriseID(enterpriseID int) (employees []models.Employee, err error) {
	return repository.GetAllEmployeesByEnterpriseID(enterpriseID)
}

func GetEmployeeByID(id int) (e models.Employee, err error) {
	return repository.GetEmployeeByID(id)
}

func CreateEmployee(e models.Employee) error {
	return repository.CreateEmployee(e)
}
