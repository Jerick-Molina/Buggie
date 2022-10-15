package server

import (
	"fmt"
	"net/http"

	"github.com/Jerick-Molina/Buggie/internal/security"
	"github.com/gin-gonic/gin"
)

// type MiddleWare struct {
// 	server *Serverconst
// }

//Checks token and sets values.
func MiddleWareToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req CreateAccountParams
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotAcceptable, "Err")
			return
		}

		fmt.Println("heelo")
		ctx.Next()
	}
}

type SignInParams struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (server *Server) SignIn(ctx *gin.Context) {

	var credits SignInParams

	if err := ctx.BindJSON(&credits); err != nil {
		//ERR
		return
	}
	credits.Password = security.HashPassword(credits.Password)

	acc, err := server.store.ValidSignInCredentials(credits.Email, credits.Password)
	if err != nil {

	}

	//token, err := security.CreateAccessToken(acc)

	ctx.JSON(http.StatusOK, acc)
}

// If it passes theres requirements then it goes to create account.
func (server *Server) CreateAccountRequirements(ctx *gin.Context) {

}

// If it passes theres requirements then it goes to create company.
func (server *Server) CreateCompanyRequirements(ctx *gin.Context) {

}
