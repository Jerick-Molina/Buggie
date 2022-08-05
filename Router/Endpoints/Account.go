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

var usr properties.User

func CreateAccount(x *gin.Context) {

	//Checks if input  meets requirements
	companyCode := x.GetHeader("CompanyCode")
	if len(companyCode) != 8 {
		x.IndentedJSON(http.StatusBadRequest, "Company code must be 8 characters long")
		return
	}
	if err := x.BindJSON(&usr); err != nil {
		x.IndentedJSON(http.StatusBadRequest, "Invalid Inputs")
		return
	}
	//Check if user already exist
	code_Find, message_Find := access.AccountFind(&usr)
	if code_Find != 200 {
		x.IndentedJSON(code_Find, message_Find)
		return
	}
	//if valid: Checks for input validation
	usr.Password = security.HashPassword(usr.Password)
	code_Create, message_Create := access.AccountCreateByCode(&usr, companyCode)
	if code_Create != 200 {

		x.IndentedJSON(code_Create, message_Create)
		return
	}
	//If oksatus then returns jwtaccess token

	token, err := security.JwtAccessToken(usr)
	if err != nil {
		x.IndentedJSON(http.StatusInternalServerError, "Something went wrong on out side")
		return
	}
	x.IndentedJSON(http.StatusAccepted, token)
}

func CreateCompany(x *gin.Context) {

	//Checks if input meets requirements
	CompanyName := x.GetHeader("CompanyName")
	if len(CompanyName) < 5 {
		x.IndentedJSON(http.StatusBadRequest, "Company Either empty or must be 5 characters or more")
		return
	}
	if err := x.BindJSON(&usr); err != nil {
		x.IndentedJSON(http.StatusBadRequest, "Invalid Inputs")
		return
	} else {
		if usr.Email == "" || usr.Password == "" {
			x.IndentedJSON(http.StatusNotAcceptable, "Invalid credidentials")
			return
		}
	}
	//If valid: Checks if user exist
	code_Exist, message_Exist := access.AccountFind(&usr)
	if code_Exist != 200 {
		x.IndentedJSON(code_Exist, message_Exist)
		return
	}
	//If valid does create companyCode
	companyId, code_Comp, message_Comp := access.CreateCompany(CompanyName)
	if code_Comp != http.StatusOK {
		x.IndentedJSON(code_Comp, message_Comp)
		return
	}
	usr.CompanyId = companyId
	usr.Password = security.HashPassword(usr.Password)
	code_Acc, message_Acc := access.AccountCreateByCompany(&usr)
	if code_Acc != http.StatusOK {

		x.IndentedJSON(code_Comp, message_Acc)
		return
	}

	token, err := security.JwtAccessToken(usr)
	if err != nil {
		x.IndentedJSON(http.StatusInternalServerError, "Something went wrong on out side")
		return
	}
	x.IndentedJSON(http.StatusAccepted, token)
}

func SignIn(x *gin.Context) {

	if err := x.BindJSON(&usr); err != nil {
		x.IndentedJSON(http.StatusBadRequest, "Invalid Inputs")
	} else {
		if usr.Email == "" || usr.Password == "" {
			x.IndentedJSON(http.StatusNotAcceptable, "Invalid credidentials")
			return
		}
	}
	usr.Password = security.HashPassword(usr.Password)
	if code, message := access.SignIn(&usr); code != http.StatusOK {
		x.IndentedJSON(code, message)
		return
	}

	token, err := security.JwtAccessToken(usr)
	if err != nil {
		x.IndentedJSON(http.StatusInternalServerError, "Something went wrong on out side")
		return
	}
	x.IndentedJSON(http.StatusAccepted, token)
}

func GetUsersRoles(x *gin.Context) {
	validRoles := []string{"Admin", "Associate"}

	claims, code, message := security.IsAuthorized(validRoles, x.GetHeader("Authorization"))
	if code != http.StatusOK {
		fmt.Println("hello")
		x.IndentedJSON(code, message)
		return
	}
	strComp := fmt.Sprintf("%v", claims["companyId"])
	strRole := fmt.Sprintf("%v", claims["Role"])
	compIdVal, err := strconv.ParseInt(strComp, 0, 32)
	if err != nil {
		fmt.Println("hello2")

		x.IndentedJSON(http.StatusUnauthorized, "Something happened in our end.")
		return
	}

	users, code, message := access.AssignmentsRole(int(compIdVal), strRole)
	if code != http.StatusOK {
		fmt.Println("hello")

		x.JSON(code, message)
		return
	}

	x.JSON(code, users)

}
