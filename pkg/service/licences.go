package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllLicencesByEnterpriseID(enterpriseID int) (licences []models.Licence, err error) {
	return repository.GetAllLicencesByEnterpriseID(enterpriseID)
}

func GetLicenceByID(id int) (l models.Licence, err error) {
	return repository.GetLicenceByID(id)
}
