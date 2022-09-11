package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Testing(ctx *gin.Context) {

	var e = "MyTest@Gmail.com"
	//server.db.TestDB("MyTest@Gmail.com")

	server.db.TestDB()
	ctx.JSON(http.StatusOK, "heelo")
}
