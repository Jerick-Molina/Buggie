package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	access "github.com/Jerick-Molina/Buggie/src/Database/Access"
	security "github.com/Jerick-Molina/Buggie/src/Security"
	"github.com/gin-gonic/gin"
)

func GetGeneralDashboard(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code, message := security.IsAuthorized(validRoles, x.GetHeader("Authorization"))

	if code != http.StatusOK {

		x.IndentedJSON(code, message)
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
	tktDta, code, message := access.TicketFindByCompany(int(compIdVal))

	if code != http.StatusOK {
		fmt.Println("2")
		x.IndentedJSON(code, message)
		return
	}
	projectDta, code, message := access.ProjectSearchByCompanyId(int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("3")
		x.IndentedJSON(code, message)
		return
	}
	usrData, code, message := access.AccountsFindByCompany(int(compIdVal))
	if code != http.StatusOK {
		fmt.Println("3")
		x.IndentedJSON(code, message)
		return
	}
	fmt.Println("1")
	fmt.Println(tktDta)
	x.IndentedJSON(http.StatusOK, []any{tktDta, projectDta, usrData})
}

func GetAssignedDataForTicket(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code, message := security.IsAuthorized(validRoles, x.GetHeader("Authorization"))

	if code != http.StatusOK {

		x.IndentedJSON(code, message)
		return
	}
	// prjtDta := access.ProjectSearchByCompanyId(claims["sss"])
	str := fmt.Sprintf("%v", claims["companyId"])

	compIdVal, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		fmt.Println("1")
		x.IndentedJSON(http.StatusInternalServerError, "Something happened in our end")
		return
	}

	projData, code, message := access.ProjectSearchByCompanyId(int(compIdVal))
	if code != http.StatusOK {
		x.JSON(code, message)
	}
	var list []any
	for _, proj := range projData {

		usr, code, message := access.AccountsFindByProject(proj.Id, int(compIdVal))
		if code != http.StatusOK {

			x.JSON(code, message)
			break

		}
		list = append(list, usr)
	}

	x.JSON(http.StatusOK, []any{projData, list})
}
