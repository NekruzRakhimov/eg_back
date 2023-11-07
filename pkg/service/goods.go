package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllGoodsByEnterpriseID(enterpriseID int) (goods []models.Good, err error) {
	return repository.GetAllGoodsByEnterpriseID(enterpriseID)
}
