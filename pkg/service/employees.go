package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllEmployeesByEnterpriseID(enterpriseID int) (employees []models.Employee, err error) {
	return repository.GetAllEmployeesByEnterpriseID(enterpriseID)
}
