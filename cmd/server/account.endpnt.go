package server

import (
	"net/http"

	db "github.com/Jerick-Molina/Buggie/internal/database"
	"github.com/gin-gonic/gin"
)

type CreateAccountParams struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

func (server *Server) createAccountByCode(ctx *gin.Context) {
	var req CreateAccountParams
	compCode := ctx.GetHeader("CompanyCode")

	if len(compCode) != 8 {
		ctx.JSON(http.StatusBadRequest, "Company Code err")
		return
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Err")
		return
	}

	// Checks if account exist
	if err := server.store.SearchEmailExistValidation(req.Email); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Err")
		return
	}

	arg := db.CreateAccByCompanyCodeParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	acc, err := server.store.CreateAccountByCompanyCode(arg)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Err")
		return
	}

	ctx.JSON(http.StatusOK, acc)
}
