package main

import (
	"github.com/Jerick-Molina/Buggie/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	database.Open()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "projectId", "content-type"},
		AllowMethods:    []string{"GET", "POST"},
	}))
	defer router.Run(":8080")

}
