package main

import (
	"test6/app/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/create", models.Create)
		v1.GET("/get/:id", models.Get)
		// v1.PUT("/update/:id", models.Update)
		v1.PUT("/update", models.Update)
		v1.DELETE("/delete/:id", models.Delete)
	}

	router.Run(":8080")
}
