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

func CreateEnterprise(c *gin.Context) {
	type bodyTemp struct {
		Body models.Enterprise `json:"body"`
	}

	var body bodyTemp

	var e models.Enterprise
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e = body.Body

	if err := service.CreateEnterprise(e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новое предприятие создано!"})
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

	structure, err := service.GetEnterpriseStructure(e.ID, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(structure) == 0 {
		structure = []models.EnterpriseStructure{}
	}

	goods, err := service.GetAllGoodsByEnterpriseID(e.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(goods) == 0 {
		goods = []models.Good{}
	}

	c.JSON(http.StatusOK, gin.H{
		"enterprise": e,
		"employees":  employees,
		"licences":   licences,
		"structure":  structure,
		"goods":      goods,
	})
}

func AddLicenceToEnterprise(c *gin.Context) {
	enterpriseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type bodyTemp struct {
		Body models.Licence `json:"body"`
	}

	var body bodyTemp

	if err = c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.AddLicenceToEnterprise(enterpriseID, body.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новая лицензия успешно добавлена"})
}
