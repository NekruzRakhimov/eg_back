package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
)

func GetBulkEnterpriseEconomicActivity(enterpriseID int) (bulkActivity []models.BulkEconomicActivity, err error) {
	years, err := repository.GetYearsOfEconomicActivity(enterpriseID)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(years); i++ {
		activity, err := repository.GetEnterpriseActivityByYear(enterpriseID, years[i])
		if err != nil {
			return nil, err
		}
		bulkActivity = append(bulkActivity, models.BulkEconomicActivity{
			Year:             years[i],
			EconomicActivity: activity,
		})
	}

	return bulkActivity, nil
}
