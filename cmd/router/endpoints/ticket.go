package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerick-Molina/Buggie/api/types"
	"github.com/Jerick-Molina/Buggie/internal/database/access"
	"github.com/Jerick-Molina/Buggie/internal/security"
	"github.com/gin-gonic/gin"
)

var tkt types.Tickets

func CreateTicket(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code := security.AccessTokenAuthorization(validRoles, x.GetHeader("Authorization"))
	if code != http.StatusOK {
		fmt.Println("ee")
		x.JSON(code, "message")
		return
	}
	if err := x.BindJSON(&tkt); err != nil {
		x.JSON(http.StatusUnauthorized, "Invalid Inputs")
		return
	}
	str := fmt.Sprintf("%v", claims["companyId"])
	compIdVal, err := strconv.ParseInt(str, 0, 32)
	if err != nil {

	}
	strs := fmt.Sprintf("%v", claims["userId"])
	userIdVal, err := strconv.ParseInt(strs, 0, 32)
	if err != nil {
	}

	if code := access.InsertTicket(int(compIdVal), int(userIdVal), &tkt); code != http.StatusOK {
		return
	}
	if tkt.AssignedTo == 0 {
		x.JSON(http.StatusOK, "Valid: No assignedTo")
		return
	}

	if code := access.QueryAccountsAssignedToTicket(int(compIdVal), tkt); code != http.StatusOK {

	}

	x.JSON(http.StatusOK, "Valid")
}

func EditTicket() {

}

func FindTicketByCompanyId() {

}

func FindTicketByProjectId() {

}
