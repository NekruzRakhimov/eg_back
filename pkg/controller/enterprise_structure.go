package controller

import (
	"eg_back/models"
	"eg_back/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEnterpriseStructureByID(c *gin.Context) {
	enterpriseID, err := strconv.Atoi(c.Param("enterprise_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	structureID, err := strconv.Atoi(c.Param("structure_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := service.GetEnterpriseStructureByID(enterpriseID, structureID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	childStructures, err := service.GetEnterpriseStructure(enterpriseID, structureID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(childStructures) == 0 {
		childStructures = []models.EnterpriseStructure{}
	}

	c.JSON(http.StatusOK, gin.H{
		"structure":        s,
		"child_structures": childStructures,
	})
}
