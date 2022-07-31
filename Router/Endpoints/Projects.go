package endpoints

import (
	"fmt"
	"net/http"

	security "github.com/Jerick-Molina/Buggie/src/Security"

	"github.com/gin-gonic/gin"
)

func ProjectCreate(x *gin.Context) {

}

func ProjectEdit(x *gin.Context) {

	validRoles := []string{"NotAdmin"}
	//token, err := security.JwtAccessToken(properties.User{Role: "Admin"})

	bToken := x.GetHeader("Authorization")
	fmt.Println(bToken[7:])

	claims, code, message := security.IsAuthorized(validRoles, bToken)
	if code != http.StatusOK {
		x.IndentedJSON(code, message)
		return
	}

	fmt.Println(claims)

}

func ProjectFind(x *gin.Context, role []string) {

	ss := x.GetHeader("Authorization")
	fmt.Println(ss)

	fmt.Println(ss)
	//	access.ProjectFindById()

	x.IndentedJSON(http.StatusOK, "Success")
}

func ProjectDelete(x *gin.Context) {

}
