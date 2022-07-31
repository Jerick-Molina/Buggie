package main

import (
	route "github.com/Jerick-Molina/Buggie/Router"
	database "github.com/Jerick-Molina/Buggie/src/Database"
	"github.com/gin-gonic/gin"
)

var routeHost string = "localhost:8080"

func main() {

	router := gin.Default()
	defer router.Run(routeHost)
	database.Open()
	route.Route(router)

}
