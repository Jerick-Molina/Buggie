package main

import (
	route "github.com/Jerick-Molina/Buggie/Router"
	database "github.com/Jerick-Molina/Buggie/src/Database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var routeHost string = "localhost:8080"

func main() {

	router := gin.Default()
	defer router.Run(routeHost)
	database.Open()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "projectId", "content-type"},
		AllowMethods:    []string{"GET", "POST"},
	}))
	route.Route(router)

}
