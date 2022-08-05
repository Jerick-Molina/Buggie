package access

import (
	"database/sql"
	"fmt"
	"net/http"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	database "github.com/Jerick-Molina/Buggie/src/Database"
)

/*
	Creates Account and returns
	httpCode, message , user
*/

func AccountCreateByCode(usr *properties.User, companyCode string) (int, string) {
	var db = database.Access()

	sqlQuery1 := "select CompanyId from Company where CompanyCode = ?"
	sqlQuery2 := "select UserId from Users where Email = ?"
	sqlExc := "insert into Users(FirstName,LastName,Email,Password,Role,CompanyId) values (?,?,?,?,?,?);"

	if err := db.QueryRow(sqlQuery1, companyCode).Scan(&usr.CompanyId); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, "Company Code is invalid "
		}
	}
	usr.Role = "Developer"
	result, err := db.Exec(sqlExc, usr.FirstName, usr.LastName, usr.Email, usr.Password, usr.Role, usr.CompanyId)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	if err := db.QueryRow(sqlQuery2, usr.Email).Scan(&usr.Id); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusInternalServerError, "Internal error"
		}
	}

	return http.StatusOK, "Ok"
}
func AccountCreateByCompany(usr *properties.User) (int, string) {
	var db = database.Access()

	sqlExc := "insert into Users(FirstName,LastName,Email,Password,Role,CompanyId) values (?,?,?,?,?,?);"
	sqlQuery := "select UserId from Users where Email = ?"
	usr.Role = "Admin"
	result, err := db.Exec(sqlExc, usr.FirstName, usr.LastName, usr.Email, usr.Password, usr.Role, usr.CompanyId)
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusInternalServerError, "something happened in out end"
	}
	fmt.Println(result)
	if err := db.QueryRow(sqlQuery, usr.Email).Scan(&usr.Id); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusInternalServerError, "Internal error"
		}
	}
	return http.StatusOK, "Ok"
}

func SignIn(usr *properties.User) (int, string) {
	var db = database.Access()

	sqlQuery := "select UserId,Email,FirstName,LastName,Role,CompanyId from Users where Email = ? and Password = ?"

	//Checks if user input matches a row, if so Authentication passes if it does not find anything that means user credidentials are wrong.
	if result := db.QueryRow(sqlQuery, usr.Email, usr.Password).Scan(&usr.Id, &usr.Email, &usr.FirstName, &usr.LastName, &usr.Role, &usr.CompanyId); result != nil {

		return http.StatusNotAcceptable, "Invalid Credidentials "
	} else {
		if result == sql.ErrConnDone {
			return http.StatusUnauthorized, "Invalid Credidentials"
		}

	}

	return http.StatusOK, "Valid user input"
}

func AccountFind(usr *properties.User) (int, string) {
	var db = database.Access()

	sqlQuery := "select Email from Users where Email = ?"

	//Checks if user input matches a row, if so Authentication passes if it does not find anything that means user credidentials are wrong.
	if result := db.QueryRow(sqlQuery, usr.Email).Scan(&usr.Email); result == nil {
		fmt.Println("ERRR?")
		return http.StatusNotAcceptable, "Invalid Credidentials"
	} else {
		if result == sql.ErrConnDone {
			return http.StatusUnauthorized, "Invalid Credidentials"
		}
	}

	return http.StatusOK, "Valid user input"
}

func AccountsFindByCompany(companyId int) ([]properties.User, int, string) {
	db := database.Access()
	sqlQuery := "select FirstName,LastName,Email,Role from Users where CompanyId = ?"
	var users []properties.User

	rows, err := db.Query(sqlQuery, companyId)
	if err != nil {
		return nil, http.StatusInternalServerError, "Query Error"
	}

	for rows.Next() {
		var usr properties.User

		if err := rows.Scan(&usr.FirstName, &usr.LastName, &usr.Email, &usr.Role); err != nil {
			return nil, http.StatusInternalServerError, "Query Error"
		}
		users = append(users, usr)
	}

	return users, http.StatusOK, "Valid"
}

func AccountsFindByProject(projectId int, companyId int) ([]properties.User, int, string) {
	db := database.Access()
	sqlQuery := `select UserId,FirstName,LastName,Email,Role from Users where UserId = any(select AssignedTo from AssignedProject where ProjectId = ? and CompanyId = ?)`
	var users []properties.User

	rows, err := db.Query(sqlQuery, projectId, companyId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusInternalServerError, "Query Error"
	}

	for rows.Next() {
		var usr properties.User

		if err := rows.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Role); err != nil {
			return nil, http.StatusInternalServerError, "Query Error"
		}
		users = append(users, usr)
	}

	return users, http.StatusOK, "Valid"
}

func AccountAssignedToProject() {

}

func AccountAssignedToTicket(companyId int, tkt properties.Tickets) (int, string) {
	db := database.Access()

	sqlExc := "insert into AssignedProject(CompanyId,ProjectId,AssignedTo) values(?,?,?)"

	results, err := db.Exec(sqlExc, companyId, tkt.ProjectId, tkt.AssignedTo)
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusInternalServerError, "something happened in out end"
	}
	fmt.Println(results)

	return http.StatusOK, "Valid"
}
