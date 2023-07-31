package controller

import (
	"eg_back/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllEnterprises(c *gin.Context) {
	e := []models.Enterprise{
		{
			Id:          1,
			Title:       "Test",
			Price:       1000,
			Description: "Test",
			Category:    "Test",
			Image:       "Test",
			Rating: struct {
				Rate  float64 `json:"rate"`
				Count int     `json:"count"`
			}{
				Rate:  5,
				Count: 2,
			},
		},
	}

	c.JSON(http.StatusOK, e)
}
