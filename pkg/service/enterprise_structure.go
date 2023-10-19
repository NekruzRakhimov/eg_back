package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetEnterpriseStructure(enterpriseID int, structureID int) (s []models.EnterpriseStructure, err error) {
	return repository.GetEnterpriseStructure(enterpriseID, structureID)
}

func GetEnterpriseStructureByID(enterpriseID int, structureID int) (s models.EnterpriseStructure, err error) {
	return repository.GetEnterpriseStructureByID(enterpriseID, structureID)
}
