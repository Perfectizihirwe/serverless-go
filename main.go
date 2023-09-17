package main

import (
	"github.com/Perfectizihirwe/serverless-go-api/initializers"
	"github.com/Perfectizihirwe/serverless-go-api/models"
	"github.com/Perfectizihirwe/serverless-go-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
	initializers.DB.AutoMigrate(&models.Post{})
}

func main() {
	app := gin.New()
	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	app.Run()
}
