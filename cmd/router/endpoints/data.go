package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerick-Molina/Buggie/internal/database/access"
	"github.com/Jerick-Molina/Buggie/internal/security"

	"github.com/gin-gonic/gin"
)

func GetGeneralDashboard(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code := security.AccessTokenAuthorization(validRoles, x.GetHeader("Authorization"))

	if code != http.StatusOK {

		x.IndentedJSON(code, "")
		return
	}
	// prjtDta := access.ProjectSearchByCompanyId(claims["sss"])
	str := fmt.Sprintf("%v", claims["companyId"])

	compIdVal, err := strconv.ParseInt(str, 0, 32)

	if err != nil {
		fmt.Println("1")
		x.IndentedJSON(http.StatusInternalServerError, "Something happened in our end")
		return
	}
	tktDta, code := access.QueryTicketByCompanyId(int(compIdVal))

	if code != http.StatusOK {
		fmt.Println("2")
		x.IndentedJSON(code, "")
		return
	}
	projectDta, code := access.QueryProjectByCompanyId(int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("3")
		x.IndentedJSON(code, "message")
		return
	}
	usrData, code := access.QueryAccountsByCompanyId(int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("3")
		x.IndentedJSON(code, "message")
		return
	}
	fmt.Println("1")
	fmt.Println(tktDta)
	x.IndentedJSON(http.StatusOK, []any{tktDta, projectDta, usrData})
}

func GetAssignedDataForTicket(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code := security.AccessTokenAuthorization(validRoles, x.GetHeader("Authorization"))

	if code != http.StatusOK {

		x.IndentedJSON(code, "")
		return
	}
	// prjtDta := access.ProjectSearchByCompanyId(claims["sss"])
	str := fmt.Sprintf("%v", claims["companyId"])

	compIdVal, err := strconv.ParseInt(str, 0, 32)

	if err != nil {
		fmt.Println("1")
		x.IndentedJSON(http.StatusInternalServerError, "Something happened in our end")
		return
	}

	projectDta, code := access.QueryProjectByCompanyId(int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("3")
		x.IndentedJSON(code, "message")
		return
	}
	var list []any
	for _, proj := range projectDta {

		usr, code := access.QueryAccountsAssignedToProject(proj.Id, int(compIdVal))
		if code != http.StatusOK {

			x.JSON(code, "message")
			break

		}
		list = append(list, usr)
	}
	x.IndentedJSON(http.StatusOK, []any{projectDta, list})
}
