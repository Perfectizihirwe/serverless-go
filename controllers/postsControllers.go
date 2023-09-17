package controllers

import (
	"net/http"

	"github.com/Perfectizihirwe/serverless-go-api/initializers"
	"github.com/Perfectizihirwe/serverless-go-api/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var Body struct {
		Body  string
		Title string
	}

	c.Bind(&Body)

	post := models.Post{Title: Body.Title, Body: Body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func DeletePost(c *gin.Context) {
	var post models.Post
	initializers.DB.Delete(&post, c.Param("id"))

	result := initializers.DB.Delete(&models.Post{}, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
