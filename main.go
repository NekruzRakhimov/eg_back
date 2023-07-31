package main

import (
	"eg_back/db"
	"eg_back/logger"
	"eg_back/routes"
	"eg_back/utils"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	db.StartDbConnection()
	routes.RunAllRoutes()
}
