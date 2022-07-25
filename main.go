package main

import (
	"github.com/gin-gonic/gin"
)

var routeHost string = "localhost:8080"

func main() {

	router := gin.Default()
	defer router.Run(routeHost)
	Route(router)
}
