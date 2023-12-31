package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllEnterprises(filter models.Filter) (e []models.Enterprise, err error) {
	return repository.GetAllEnterprises(filter)
}

func GetEnterprisesCount(filter models.Filter) (count int, err error) {
	return repository.GetEnterprisesCount(filter)
}

func GetEnterpriseByID(id int) (e models.Enterprise, err error) {
	return repository.GetEnterpriseByID(id)
}

func CreateEnterprise(e models.Enterprise) error {
	return repository.CreateEnterprise(e)
}

func AddLicenceToEnterprise(enterpriseID int, licence models.Licence) error {
	return repository.AddLicenceToEnterprise(enterpriseID, licence)
}
