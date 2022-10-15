package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/Jerick-Molina/Buggie/internal/access/database"
// 	"github.com/Jerick-Molina/Buggie/internal/security"
// 	"github.com/gin-gonic/gin"
// )

// type Database struct {
// 	*sql.DB
// }

// func Access(db *sql.DB) *Database {

// 	return &Database{db}
// }

// func (db *Engine) SignIn(ctx *gin.Context) {

// 	if err := ctx.BindJSON(&usr); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid Inputs")
// 	} else {
// 		if usr.Email == "" || usr.Password == "" {

// 			return
// 		}
// 	}
// 	usr.Password = security.HashPassword(usr.Password)
// 	fmt.Println("heee")
// 	if code := database.QueryCredidentials(); code != 200 {
// 		ctx.JSON(code, "help")
// 		return
// 	}

// 	fmt.Printf("Hello!")
// 	token, err := security.CreateAccessToken(usr)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, "Something went wrong on out side")
// 		return
// 	}
// 	ctx.JSON(http.StatusAccepted, token)
// }
