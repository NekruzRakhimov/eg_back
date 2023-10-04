package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetAllClassifiers(filter models.Filter) (classifiers []models.Classifier, err error) {
	return repository.GetAllClassifiers(filter)
}

func GetClassifierByCode(code string) (c models.Classifier, err error) {
	return repository.GetClassifierByCode(code)
}

func GetAllClassifiersCount(filter models.Filter) (count int, err error) {
	return repository.GetAllClassifiersCount(filter)
}

func GetClassifierItems(classifierID, ItemID int) (items []models.ClassifierItem, err error) {
	return repository.GetClassifierItems(classifierID, ItemID)
}
