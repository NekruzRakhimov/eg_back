package controller

import (
	"eg_back/models"
	"eg_back/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIndustrialObjectsKeyIndicators(c *gin.Context) {
	var filter models.ReportFilter
	filter.DateFrom = c.Query("_date_from")
	filter.DateTo = c.Query("_date_to")
	filter.Location = c.Query("_location")
	filter.EconomicActivity = c.Query("_economic_activity")
	filter.OrganizationAge, _ = strconv.Atoi(c.Query("_organization_age"))
	filter.YearlyTurnover, _ = strconv.Atoi(c.Query("_yearly_turnover"))
	filter.AuthorizedCapital, _ = strconv.Atoi(c.Query("_authorized_capital"))

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
	fmt.Printf("%#v", filter)

	r, err := service.GetReports(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}
