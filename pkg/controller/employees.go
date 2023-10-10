package controller

import (
	"eg_back/models"
	"eg_back/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, err := service.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func CreateEmployee(c *gin.Context) {
	enterpriseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	type bodyTemp struct {
		Body models.Employee `json:"body"`
	}

	var body bodyTemp
	if err = c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	body.Body.EnterpriseID = enterpriseID
	if err = service.CreateEmployee(body.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "сотрудник успешно создан!"})
}
