package endpoints

import (
	"fmt"
	"net/http"

	security "github.com/Jerick-Molina/Buggie/src/Security"
	"github.com/gin-gonic/gin"
)

func GetGeneralDashboard(x *gin.Context) {
	validRoles := []string{"Admin", "Manager"}

	claims, code, message := security.IsAuthorized(validRoles, x.GetHeader("Authorize"))

	if code != http.StatusOK {
		x.IndentedJSON(code, message)
	}

	fmt.Println(claims)

}

func GetProjectDashboard() {

}
