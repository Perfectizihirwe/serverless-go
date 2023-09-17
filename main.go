package main

import (
	"context"
	"log"

	"github.com/Perfectizihirwe/serverless-go-api/initializers"
	"github.com/Perfectizihirwe/serverless-go-api/models"
	"github.com/Perfectizihirwe/serverless-go-api/routes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
	initializers.DB.AutoMigrate(&models.Post{})
	app := gin.Default()
	router := app.Group("/api/v1")
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.AddRoutes(router)

	ginLambda = ginadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
