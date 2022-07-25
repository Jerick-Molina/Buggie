package main

import (
	"github.com/gin-gonic/gin"
)

var route *gin.Engine
var api string = "/API"

func Route(x *gin.Engine) {
	//Since pointers were used we only need to declare it once. Saving memory.
	route = x
	//Each type of routes
	Project()
	Ticket()
	Account()
	Dashboard()
}
func GET(rs string, handler gin.HandlerFunc) {
	route.GET(api+rs, handler)
}

func POST(routes string, handler gin.HandlerFunc) {
	route.POST(api+routes, handler)
}
func Account() {
	//GET

	//POST
	POST("/Create/User", NoImpletmentedException)
	POST("/Create/Companyt", NoImpletmentedException)
	POST("/SignIn", NoImpletmentedException)
}
func Project() {
	//GET

	//POST
	GET("/Project/Create", NoImpletmentedException)
}

func Ticket() {
	//GET

	//POST
}

func Dashboard() {
	//GET
	GET("/Dashboard", NoImpletmentedException)
	//POST
}
