package access

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/Jerick-Molina/Buggie/api/types"
// )

// type DbAccess struct {
// 	db *sql.DB
// }

// func InsertAccountByCode(usr *types.User, companyCode string) int {

// 	sqlQuery1 := "select CompanyId from Company where CompanyCode = ?"
// 	sqlQuery2 := "select UserId from Users where Email = ?"
// 	sqlExc := "insert into Users(FirstName,LastName,Email,Password,Role,CompanyId) values (?,?,?,?,?,?);"

// 	if err := database.Db.QueryRow(sqlQuery1, companyCode).Scan(&usr.CompanyId); err != nil {
// 		if err == sql.ErrNoRows {
// 			return http.StatusBadRequest
// 		}
// 	}
// 	//
// 	usr.Role = "Developer"
// 	result, err := database.Db.Exec(sqlExc, usr.FirstName, usr.LastName, usr.Email, usr.Password, usr.Role, usr.CompanyId)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(result)
// 	if err := database.Db.QueryRow(sqlQuery2, usr.Email).Scan(&usr.Id); err != nil {
// 		if err == sql.ErrNoRows {
// 			return http.StatusInternalServerError
// 		}
// 	}

// 	return http.StatusOK
// }

// func InsertAccountByCompany(usr *types.User) int {

// 	sqlExc := "insert into Users(FirstName,LastName,Email,Password,Role,CompanyId) values (?,?,?,?,?,?)"

// 	usr.Role = "Admin"

// 	result, err := db.Exec(sqlExc, usr.FirstName, usr.LastName, usr.Email, usr.Password,
// 		usr.Role, usr.CompanyId)

// 	if err != nil {
// 		return http.StatusInternalServerError
// 	}
// 	fmt.Println(result)
// 	return http.StatusOK
// }

// func QueryAccountsAssignedToProject(projectId int, companyId int) ([]types.User, int) {

// 	var usrs []types.User
// 	sqlQuery := `select UserId,FirstName,LastName,Email,Role from Users where UserId = any(select AssignedTo from AssignedProject where ProjectId = ? and CompanyId = ?)`

// 	rows, err := db.Query(sqlQuery, projectId, companyId)
// 	if err != nil {
// 		return nil, http.StatusInternalServerError
// 	}

// 	for rows.Next() {
// 		var usr types.User

// 		if err := rows.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Role); err != nil {
// 			return nil, http.StatusInternalServerError
// 		}
// 		usrs = append(usrs, usr)
// 	}
// 	return usrs, http.StatusOK

// }

// func QueryAccountsAssignedToTicket(companyId int, tkt types.Tickets) int {
// 	sqlExc := "insert into AssignedProject(CompanyId,ProjectId,AssignedTo) values(?,?,?)"

// 	results, err := db.Exec(sqlExc, companyId, tkt.ProjectId, tkt.AssignedTo)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return http.StatusInternalServerError
// 	}
// 	fmt.Println(results)

// 	return 200
// }

// func QueryAccountsByCompanyId(compId int) ([]types.User, int) {

// 	return nil, 200
// }
// func (db *DbAccess) QueryCredidentials(usr *types.User) int {

// 	return http.StatusOK
// }

// func QueryAccount(usr *types.User) int {

// 	sqlQuery := "select Email from Users where Email = ?"

// 	if result := db.QueryRow(sqlQuery, usr.Email).Scan(&usr.Email); result == nil {
// 		return http.StatusNotAcceptable
// 	} else {
// 		if result == sql.ErrConnDone {
// 			return http.StatusUnauthorized
// 		}
// 	}

// 	return http.StatusOK
// }
