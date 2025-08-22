package handlers

import (
	"blog-app/V0/config"
	"blog-app/V0/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	post := models.Post{
		Title:   c.PostForm("title"),
		Content: c.PostForm("content"),
	}

	result := config.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, post) // return the created post with ID, timestamps
}

func GetPost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get End point"})
}
