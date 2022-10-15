package server

import (
	db "github.com/Jerick-Molina/Buggie/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	store           *db.Store
	Status          map[string]string
	Error           string
	AllowAllOrigins bool
	AllowedHeaders  []string
	AllowedMethods  []string
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "projectId", "content-type"},
		AllowMethods:    []string{"GET", "POST"},
	}))

	server.router = route
	server.EndpointsHandler()
	defer route.Run("localhost:8080")
	return server
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func (servers *Server) EndpointsHandler() {
	//This is where you init your routes

	api := servers.router.Group("/api")
	{

		auth_Api := api.Group("/auth")
		{
			auth_Api.POST("/SignIn")
		}
		api.POST("/Create_Account", servers.CreateAccountRequirements, servers.createAccountByCode)
		api.POST("/Create_Company", servers.CreateCompany)

	}
}
