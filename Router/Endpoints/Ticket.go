package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	access "github.com/Jerick-Molina/Buggie/src/Database/Access"
	security "github.com/Jerick-Molina/Buggie/src/Security"
	"github.com/gin-gonic/gin"
)

var tkt properties.Tickets

func TicketCreate(x *gin.Context) {
	validRoles := []string{"Admin", "Developer"}

	claims, code, message := security.IsAuthorized(validRoles, x.GetHeader("Authorization"))
	if code != http.StatusOK {
		fmt.Println("ee")
		x.JSON(code, message)
		return
	}
	if err := x.BindJSON(&tkt); err != nil {
		x.JSON(http.StatusUnauthorized, "Invalid Inputs")
		return
	}
	fmt.Println(tkt)
	str := fmt.Sprintf("%v", claims["companyId"])
	compIdVal, err := strconv.ParseInt(str, 0, 32)
	if err != nil {

	}
	strs := fmt.Sprintf("%v", claims["userId"])
	userIdVal, err := strconv.ParseInt(strs, 0, 32)
	if err != nil {

	}
	code, create_message := access.CreateTicket(int(compIdVal), int(userIdVal), &tkt)

	if code != http.StatusOK {
		x.JSON(code, create_message)
		return
	}
	if tkt.AssignedTo == 0 {
		x.JSON(http.StatusOK, "Valid: No assignedTo")
		return
	}
	code, acc_message := access.AccountAssignedToTicket(int(compIdVal), tkt)
	if code != http.StatusOK {
		x.JSON(code, acc_message)
		return
	}
	x.JSON(http.StatusOK, "Valid")
}

// func TicketFindByCompany(companyId int) ([]properties.Tickets, int, string) {

// }
// func TicketsFindByProject(projectId int, companyId int) ([]properties.Tickets, int, string) {

// }

func TicketEdit() {

}
