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
	//e := []models.Enterprise{
	//	{
	//		Id:          1,
	//		Title:       "Test",
	//		Price:       1000,
	//		Description: "Test",
	//		Category:    "Test",
	//		Image:       "Test",
	//		//Rating: struct {
	//		//	Rate  float64 `json:"rate"`
	//		//	Count int     `json:"count"`
	//		//}{
	//		//	Rate:  5,
	//		//	Count: 2,
	//		//},
	//	},
	//}
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

	c.JSON(http.StatusOK, e)
}
