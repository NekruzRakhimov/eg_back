package controller

import (
	"eg_back/models"
	"eg_back/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllClassifiers(c *gin.Context) {
	var filter models.Filter
	filter.Query = c.Query("_query")
	filter.Sort = c.Query("_sort")

	if c.Query("_limit") != "" {
		filter.Limit, _ = strconv.Atoi(c.Query("_limit"))
	}

	if c.Query("_page") != "" {
		filter.Page, _ = strconv.Atoi(c.Query("_page"))
	}

	fmt.Printf("%+v", filter)

	count, err := service.GetAllClassifiersCount(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.Writer.Header().Set("x-total-count", strconv.Itoa(count))
	classifiers, err := service.GetAllClassifiers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if len(classifiers) == 0 {
		classifiers = []models.Classifier{}
	}

	c.JSON(http.StatusOK, classifiers)
}
