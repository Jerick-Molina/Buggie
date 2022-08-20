package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerick-Molina/Buggie/internal/database/access"
	"github.com/Jerick-Molina/Buggie/internal/security"
	"github.com/gin-gonic/gin"
)

func CreateProject() {

}

func EditProject() {

}

func SearchProjectByCompanyId() {

}

func SearchProjectBySelection(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code := security.AccessTokenAuthorization(validRoles, x.GetHeader("Authorization"))
	if code != http.StatusOK {

		return
	}
	str := fmt.Sprintf("%v", claims["companyId"])
	compIdVal, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		//x.IndentedJSON(http.StatusInternalServerError, "Something happened in our end.")
		return
	}
	projId := x.GetHeader("projectId")
	projIdVal, err := strconv.ParseInt(projId, 0, 32)

	if err != nil {
		//	x.IndentedJSON(http.StatusBadRequest, "Invalid input Input")
		return
	}

	tktData, code := access.QueryProjectByProjectId(int(compIdVal), int(projIdVal))
	if code != http.StatusOK {
		return
	}
	usrData, code := access.QueryAccountsAssignedToProject(int(compIdVal), int(projIdVal))
	if code != http.StatusOK {
		return
	}
	projectData, code := access.QueryProjectByProjectId(int(compIdVal), int(projIdVal))
	if code != http.StatusOK {
		return
	}
	x.JSON(http.StatusOK, []any{tktData, usrData, projectData})
}
