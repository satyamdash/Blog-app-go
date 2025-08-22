package main

import (
	"blog-app/V0/config"
	"blog-app/V0/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	r.GET("/", handlers.GetPost)

	r.POST("/post", handlers.CreatePost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(r.Run(":" + port))
}
