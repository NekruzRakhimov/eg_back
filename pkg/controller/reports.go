package controller

import (
	"eg_back/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndustrialObjectsKeyIndicators(c *gin.Context) {
	//r := models.Report{
	//	Active:         1,
	//	Liquidated:     2,
	//	Stopped:        3,
	//	OnBankruptcy:   4,
	//	LeasedOut:      5,
	//	ReOrganized:    6,
	//	CustomsDebtors: 7,
	//	TaxDebtors:     8,
	//}

	r, err := service.GetReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}
