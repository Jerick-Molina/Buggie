package access

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/Jerick-Molina/Buggie/api/types"
// )

// var roles = map[int]string{0: "Admin", 1: "Associate", 2: "Developer"}

// func InsertAssignedToTicket() {

// }

// func InsertAssignedToProject() {

// }
// func QueryAssignedRoles(compId int, usrRole string) ([]types.User, int) {

// 	var usrs []types.User

// 	sqlQuery := "select FirstName,LastName,Email,Role from Users where CompanyId = ? and Role = ?"

// 	roleIndex := roleFilter(usrRole)
// 	for i := roleIndex + 1; i < len(roles); i++ {

// 		rows, err := db.Query(sqlQuery, compId, roles[i])
// 		fmt.Println(roles[i])
// 		if err != nil {
// 			fmt.Println(err.Error())

// 			return nil, http.StatusInternalServerError
// 		}

// 		for rows.Next() {
// 			var usr types.User
// 			if err := rows.Scan(&usr.FirstName, &usr.LastName, &usr.Email, &usr.Role); err != nil {

// 				fmt.Println(err.Error())
// 				return nil, http.StatusInternalServerError
// 			}

// 			usrs = append(usrs, usr)
// 		}
// 	}
// 	return []types.User{}, 200
// }

// func roleFilter(usrRole string) int {
// 	var roleIndex = 0
// 	for s, n := range roles {
// 		if n == usrRole {
// 			roleIndex = s
// 		}
// 	}
// 	return roleIndex
// }
