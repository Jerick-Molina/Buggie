package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	access "github.com/Jerick-Molina/Buggie/src/Database/Access"
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

func ProjectSearchByCompany(x *gin.Context) {

	ss := x.GetHeader("Authorization")
	fmt.Println(ss)

	fmt.Println(ss)
	//	access.ProjectFindById()

	x.IndentedJSON(http.StatusOK, "Success")
}

func ProjectSearchBySelection(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code, message := security.IsAuthorized(validRoles, x.GetHeader("Authorization"))
	if code != http.StatusOK {
		fmt.Println("ee")
		x.IndentedJSON(code, message)
		return
	}
	str := fmt.Sprintf("%v", claims["companyId"])
	compIdVal, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		x.IndentedJSON(http.StatusInternalServerError, "Something happened in our end.")
		return
	}
	projId := x.GetHeader("projectId")
	projIdVal, err := strconv.ParseInt(projId, 0, 32)

	if err != nil {
		x.IndentedJSON(http.StatusBadRequest, "Invalid input Input")
		return
	}
	tktData, code, message := access.TicketsFindByProject(int(projIdVal), int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("ee")
		x.IndentedJSON(code, message)
		return
	}
	usrData, code, message := access.AccountsFindByProject(int(projIdVal), int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("ee")
		x.IndentedJSON(code, message)
		return
	}
	prjtData, code, message := access.ProjectSearchByProjectId(int(projIdVal), int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("ee")
		x.IndentedJSON(code, message)
		return
	}
	x.JSON(http.StatusOK, []any{tktData, usrData, prjtData})
	return
}
