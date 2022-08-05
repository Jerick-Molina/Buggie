package router

import (
	"net/http"

	endpoints "github.com/Jerick-Molina/Buggie/Router/Endpoints"
	"github.com/gin-gonic/gin"
)

var secret []byte
var router *gin.Engine

// var Roles [4]string  {"Admin","Associate","Developer"}
var api string = "/API"

func Route(x *gin.Engine) {
	router = x

	//Initialize routes
	Account()
	Ticket()
	Data()
	Project()
	Assigned()
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
	GET("/Assign/Roles", endpoints.CreateAccount)
	GET("/Assign/Projects", endpoints.CreateAccount)
	//POST
	POST("/Create/User", endpoints.CreateAccount)
	POST("/Create/Companyt", endpoints.CreateCompany)
	POST("/SignIn", endpoints.SignIn)
	POST("/SignIns", endpoints.GetGeneralDashboard)
}
func Project() {
	//GET
	GET("/Project", endpoints.ProjectSearchBySelection)
	GET("/Project/Create", NoImpletmentedException)

	//POST
}

func Assigned() {
	//Get
	GET("/Assigned/Roles", endpoints.GetUsersRoles)

	//post
}
func Ticket() {
	//GET

	//POST
	POST("/Tickets/Create", endpoints.TicketCreate)

}

func Data() {
	//GET
	GET("/Dashboard", endpoints.GetGeneralDashboard)
	GET("/Tickets/Create", endpoints.GetAssignedDataForTicket)
	//POST
}

func NoImpletmentedException(x *gin.Context) {

	x.IndentedJSON(http.StatusNotImplemented, "Not implemented")
}
