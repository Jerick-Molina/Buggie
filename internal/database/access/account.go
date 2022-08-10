package access

import (
	"database/sql"
	"fmt"
	"net/http"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	"github.com/Jerick-Molina/Buggie/internal/database"
)

func AccountCreateByCode(usr *properties.User, companyCode string) (int, string) {

	sqlQuery1 := "select CompanyId from Company where CompanyCode = ?"
	sqlQuery2 := "select UserId from Users where Email = ?"
	sqlExc := "insert into Users(FirstName,LastName,Email,Password,Role,CompanyId) values (?,?,?,?,?,?);"

	if err := database.Db.QueryRow(sqlQuery1, companyCode).Scan(&usr.CompanyId); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, "Company Code is invalid "
		}
	}
	//
	usr.Role = "Developer"
	result, err := database.Db.Exec(sqlExc, usr.FirstName, usr.LastName, usr.Email, usr.Password, usr.Role, usr.CompanyId)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	if err := database.Db.QueryRow(sqlQuery2, usr.Email).Scan(&usr.Id); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusInternalServerError, "Internal error"
		}
	}

	return http.StatusOK, "Ok"
}
