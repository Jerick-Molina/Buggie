package server

import (
	"Buggie_2.0/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	eng *gin.Engine
	db  *database.DatabaseServer
}

func Run(_db *database.DatabaseServer) error {
	server := Server{db: _db}
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "projectId", "content-type"},
		AllowMethods:    []string{"GET", "POST"},
	}))

	server.eng = route
	server.EndpointsHandler()
	defer route.Run("localhost:8080")
	return nil
}

func (servers *Server) EndpointsHandler() {
	//This is where you init your routes

	api := servers.eng.Group("/api")
	{
		api.GET("/account", servers.Testing)
		api.GET("/accounts", servers.Testing)
	}
}
