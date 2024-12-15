package main

import (
	"pr-api/initializers"
	"pr-api/middlewares"
	"pr-api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())
	routes.SetupRoutes(router)
	router.Run()
}