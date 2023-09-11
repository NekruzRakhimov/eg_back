package routes

import (
	"eg_back/pkg/controller"
	"eg_back/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "*"},
		AllowHeaders:     []string{"Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowAllOrigins:  true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = utils.AppSettings.AppParams.PortRun
	}
	_ = r.Run(fmt.Sprintf(":%s", port))
}

func initAllRoutes(r *gin.Engine) {
	r.GET("/", PingPong)
	r.GET("/enterprises", controller.GetAllEnterprises)
	r.GET("/enterprises/:id", controller.GetEnterpriseByID)
	r.GET("/employees/:id", controller.GetEmployeeByID)
	r.GET("/licences/:id", controller.GetLicenceByID)

	r.GET("/classifiers", controller.GetAllClassifiers)
	r.GET("/classifiers/:id/items/:item_id", controller.GetClassifierItems)
	r.GET("/classifiers/:id/items", controller.GetClassifierItems)

	r.GET("/reports/technical_economic/industrial_activity_objects/key_indicators", controller.GetIndustrialObjectsKeyIndicators)
}

// PingPong Проверка
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "server is up and running"})
}
