package main

import (
	"net/http"

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
func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
