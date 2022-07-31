package router

import (
	"net/http"

	endpoints "github.com/Jerick-Molina/Buggie/Router/Endpoints"
	"github.com/gin-gonic/gin"
)

var secret []byte
var router *gin.Engine

var api string = "/API"

func Route(x *gin.Engine) {
	router = x

	//Initialize routes
	Account()
	Ticket()
	Dashboard()
	Project()
}

// func isAuthorized(isAuthorized bool, roles []string, endpoint func(x *gin.Context)) (gin.HandlerFunc, bool, *jwt.MapClaims) {

// 	if isAuthorized != true {
// 		isValid := false
// 		var claims *jwt.MapClaims
// 		return gin.HandlerFunc(func(ctx *gin.Context) {
// 			bToken := ctx.GetHeader("Authentication")
// 			if bToken != "" {
// 				token, err := jwt.Parse(bToken, func(t *jwt.Token) (interface{}, error) {
// 					if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 						return nil, nil
// 					}
// 					return secret, nil
// 				})
// 				if err != nil {

// 				}
// 				if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 					claims = &claim
// 				} else {
// 					fmt.Println("ERROR!")
// 				}
// 			}
// 		}), isValid, claims
// 	}

// 	return gin.HandlerFunc(func(ctx *gin.Context) {}), true, nil
// }

//This is to authorize

func GET(rts string, handler gin.HandlerFunc) {

	router.GET(api+rts, handler)
}

func POST(rts string, handler gin.HandlerFunc) {
	router.POST(api+rts, handler)
}

func Account() {
	//GET

	//POST
	POST("/Create/User", endpoints.CreateAccount)
	POST("/Create/Companyt", endpoints.CreateCompany)
	POST("/SignIn", endpoints.SignIn)
	POST("/SignIns", NoImpletmentedException)
}
func Project() {
	//GET

	//POST
	GET("/Project/Create", NoImpletmentedException)
}

func Ticket() {
	//GET

	//POST
}

func Dashboard() {
	//GET
	GET("/Dashboard", NoImpletmentedException)
	//POST
}

func NoImpletmentedException(x *gin.Context) {

	x.IndentedJSON(http.StatusNotImplemented, "Not implemented")
}
