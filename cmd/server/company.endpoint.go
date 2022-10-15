package server

import (
	"fmt"
	"net/http"

	db "github.com/Jerick-Molina/Buggie/internal/database"
	"github.com/Jerick-Molina/Buggie/internal/security"
	"github.com/gin-gonic/gin"
)

type CreateAccByCompanyParams struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	CompanyName string `json:"CompanyName"`
}

func (server *Server) CreateCompany(ctx *gin.Context) {

	var params CreateAccByCompanyParams
	var acc db.CreateAccByCompanyCodeParams
	if err := ctx.BindJSON(&params); err != nil {
		//ERR

		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}

	if err := server.store.SearchEmailExistValidation(params.Email); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "EXIST")
		return
	}
	comp_id, err := server.store.CreateCompany(params.CompanyName)

	if err != nil {
		//ERR
		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}
	acc.FirstName = params.FirstName
	acc.LastName = params.LastName
	acc.Email = params.Email
	acc.Password = params.Password
	acc.CompanyId = comp_id

	a, err := server.store.CreateAccountByCompany(acc)
	fmt.Println(a)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}
	token, err := security.CreateAccessToken(a.Id, a.Role, a.CompanyId)

	if err != nil {
		//ERR
		ctx.JSON(http.StatusNotAcceptable, "POopy pants3")
		return
	}

	ctx.JSON(http.StatusOK, token)
}
