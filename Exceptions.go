package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoImpletmentedException(x *gin.Context) {

	x.IndentedJSON(http.StatusNotImplemented, "Not implemented")
}

func UnAuthorizedResults(x *gin.Context) {
	x.IndentedJSON(http.StatusUnauthorized, "Unauthorized")
}
