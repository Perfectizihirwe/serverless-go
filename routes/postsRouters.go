package routes

import (
	"github.com/Perfectizihirwe/serverless-go-api/controllers"
	"github.com/gin-gonic/gin"
)

func postsRoutes(superRoute *gin.RouterGroup) {
	postsRoute := superRoute.Group("/posts")
	{
		postsRoute.POST("/", controllers.CreatePost)
		postsRoute.GET("/", controllers.GetPosts)
		postsRoute.GET("/:id", controllers.GetSinglePost)
		postsRoute.PUT("/:id", controllers.UpdatePost)
		postsRoute.DELETE("/:id", controllers.DeletePost)
	}
}
