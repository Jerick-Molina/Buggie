package endpoint

import (
	"github.com/Jerick-Molina/Buggie/api/types"
)

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/Jerick-Molina/Buggie/api/types"
// 	"github.com/Jerick-Molina/Buggie/internal/database/access"
// 	"github.com/Jerick-Molina/Buggie/internal/security"
// 	"github.com/gin-gonic/gin"
// )

var usr types.User

// func CreateAccount(x *gin.Context) {

// 	compCode := x.GetHeader("CompanyCode")

// 	if len(compCode) != 8 {
// 		x.IndentedJSON(api.Response(http.StatusNotAcceptable))
// 		return
// 	}
// 	if err := x.BindJSON(&usr); err != nil {
// 		x.IndentedJSON(api.Response(http.StatusNotAcceptable))
// 		return
// 	}

// 	if code := access.QueryAccount(&usr); code != http.StatusOK {
// 		return
// 	}
// 	usr.Password = security.HashPassword(usr.Password)
// 	if code := access.InsertAccountByCode(&usr, compCode); code != http.StatusOK {
// 		return
// 	}

// 	token, err := security.CreateAccessToken(usr)
// 	if err != nil {
// 		return
// 	}
// 	x.JSON(http.StatusOK, token)
// }

// func CreateCompany(x *gin.Context) {

// 	CompanyName := x.GetHeader("Company_Name")
// 	if len(CompanyName) != 5 {
// 		//
// 		return
// 	}
// 	if err := x.BindJSON(&usr); err != nil {
// 		//x.IndentedJSON(http.StatusBadRequest, "Invalid Inputs")
// 		return
// 	} else {
// 		if usr.Email == "" || usr.Password == "" {
// 			//	x.IndentedJSON(http.StatusNotAcceptable, "Invalid credidentials")
// 			return
// 		}
// 	}

// 	if code := access.QueryAccount(&usr); code != http.StatusOK {

// 		return
// 	}

// 	companyId, code := access.InsertCompany(CompanyName)
// 	if code != http.StatusOK {

// 		return
// 	}

// 	usr.CompanyId = companyId
// 	usr.Password = security.HashPassword(usr.Password)
// 	if code := access.InsertAccountByCompany(&usr); code != http.StatusOK {
// 		return
// 	}
// 	token, err := security.CreateAccessToken(usr)
// 	if err != nil {
// 		return
// 	}

// 	x.IndentedJSON(http.StatusOK, token)
// }

// func Importings(db *sql.DB) Access {

// 	return Access{db}
// }
// func (acs Access) SignIns() {
// 	ctx := acs.ctx
// 	if err := ctx.BindJSON(&usr); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid Inputs")
// 	} else {
// 		if usr.Email == "" || usr.Password == "" {
// 			ctx.JSON(http.StatusNotAcceptable, "Invalid credidentials")
// 			return
// 		}
// 	}
// 	usr.Password = security.HashPassword(usr.Password)
// 	fmt.Println("heee")
// 	if code := database.QueryCredidentials(acs.db, usr); code != 200 {
// 		ctx.JSON(code, "help")
// 		return
// 	}
// 	fmt.Printf("Hello!")
// 	token, err := security.CreateAccessToken(usr)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, "Something went wrong on out side")
// 		return
// 	}
// 	ctx.JSON(http.StatusAccepted, token)
// }

// func AchieveUserRole(x *gin.Context) {
// 	validRoles := []string{"Admin", "Associate"}

// 	claims, code := security.AccessTokenAuthorization(validRoles, x.GetHeader("Authorization"))
// 	if code != http.StatusOK {
// 		fmt.Println("hello")
// 		x.IndentedJSON(code, "")
// 		return
// 	}
// 	strComp := fmt.Sprintf("%v", claims["companyId"])
// 	strRole := fmt.Sprintf("%v", claims["Role"])
// 	compIdVal, err := strconv.ParseInt(strComp, 0, 32)
// 	if err != nil {
// 		fmt.Println("hello2")

// 		x.IndentedJSON(http.StatusUnauthorized, "Something happened in our end.")
// 		return
// 	}

// 	users, code := access.QueryAssignedRoles(int(compIdVal), strRole)
// 	if code != http.StatusOK {

// 	}
// 	if code != http.StatusOK {

// 		x.JSON(code, "message")
// 		return
// 	}

// 	x.JSON(code, users)

// }
