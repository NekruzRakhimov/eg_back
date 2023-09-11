package controller

import (
	"eg_back/models"
	"eg_back/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllEnterprises(c *gin.Context) {
	var filter models.Filter
	filter.Query = c.Query("_query")
	filter.Sort = c.Query("_sort")
	filter.Location = c.Query("_location_filter")

	if c.Query("_limit") != "" {
		filter.Limit, _ = strconv.Atoi(c.Query("_limit"))
	}

	if c.Query("_authorized_capital_filter") != "" {
		filter.AuthorizedCapitalFilter, _ = strconv.Atoi(c.Query("_authorized_capital_filter"))
	}

	if c.Query("_page") != "" {
		filter.Page, _ = strconv.Atoi(c.Query("_page"))
	}

	filter.EnterpriseAgeFilter = c.Query("_enterprise_age_filter")

	fmt.Printf("%+v", filter)

	count, err := service.GetEnterprisesCount(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Header().Set("x-total-count", strconv.Itoa(count))

	e, err := service.GetAllEnterprises(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(e) == 0 {
		e = []models.Enterprise{}
	}
	c.JSON(http.StatusOK, e)
}

func GetEnterpriseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, err := service.GetEnterpriseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	employees, err := service.GetAllEmployeesByEnterpriseID(e.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(employees) == 0 {
		employees = []models.Employee{}
	}

	licences, err := service.GetAllLicencesByEnterpriseID(e.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(licences) == 0 {
		licences = []models.Licence{}
	}

	c.JSON(http.StatusOK, gin.H{
		"enterprise": e,
		"employees":  employees,
		"licences":   licences,
	})
}
