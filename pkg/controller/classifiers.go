package controller

import (
	"eg_back/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllClassifiers(c *gin.Context) {
	classifier, err := service.GetAllClassifiers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classifier)
}
