package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllEnterprises() (e []models.Enterprise, err error) {
	return repository.GetAllEnterprises()
}
