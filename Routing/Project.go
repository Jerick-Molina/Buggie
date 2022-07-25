package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_Projects(x *gin.Engine) {

	x.GET("/Help", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, "Hello") })
}

func POST_Projects(x *gin.Engine) {

	x.GET("/Help", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, "Hello!") })
}
