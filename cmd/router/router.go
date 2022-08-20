package router

import "github.com/gin-gonic/gin"

func InitRoutes(x *gin.Engine) {

	api := x.Group("/Api")
	{
		api.POST("/SignIn")

		roles := api.Group("/Roles")
		{
			roles.POST("/Assign")
			roles.POST("/Assigned")
		}
		account := api.Group("/User")
		{
			account.POST("/Create")
			account.POST("/Create/Company")
		}
		project := api.Group("/Project")
		{
			//GET
			project.GET("/Assign")
			project.POST("/Create")
		}
		ticket := api.Group("/Ticket")
		{
			//GET
			ticket.POST("/Create")
			//POST
		}
		data := api.Group("/Data")
		{
			//GET
			data.GET("/Dashboard")
			//POST
		}

	}
}
