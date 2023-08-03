package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllClassifiers(filter models.Filter) (classifiers []models.Classifier, err error) {
	return repository.GetAllClassifiers(filter)
}

func GetAllClassifiersCount(filter models.Filter) (count int, err error) {
	return repository.GetAllClassifiersCount(filter)
}
