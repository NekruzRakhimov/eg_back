package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllClassifiers() (classifiers []models.Classifier, err error) {
	return repository.GetAllClassifiers()
}
