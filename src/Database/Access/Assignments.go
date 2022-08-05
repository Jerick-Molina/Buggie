package access

import (
	"fmt"
	"net/http"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	database "github.com/Jerick-Molina/Buggie/src/Database"
)

var roles = map[int]string{0: "Admin", 1: "Associate", 2: "Developer"}

//Role filter: Only Admin can edit any below, Associate cannot edit other associate roles only developers

func AssignmentsRole(compId int, usrRole string) ([]properties.User, int, string) {
	db := database.Access()
	var users []properties.User
	sqlQuery := "select FirstName,LastName,Email,Role from Users where CompanyId = ? and Role = ?"

	roleIndex := roleFilter(usrRole)

	for i := roleIndex + 1; i < len(roles); i++ {

		rows, err := db.Query(sqlQuery, compId, roles[i])
		fmt.Println(roles[i])
		if err != nil {
			fmt.Println(err.Error())

			return nil, http.StatusInternalServerError, err.Error()
		}

		for rows.Next() {
			var usr properties.User
			if err := rows.Scan(&usr.FirstName, &usr.LastName, &usr.Email, &usr.Role); err != nil {

				fmt.Println(err.Error())
				return nil, http.StatusInternalServerError, err.Error()
			}

			users = append(users, usr)
		}
	}

	return users, http.StatusOK, "valid"
}

//Role filter: Only Admin can edit any below, Associate cannot edit other associate roles only developers
func AssignmentsProject(compId int, role string) {

}

func roleFilter(usrRole string) int {

	var roleIndex = 0
	for s, n := range roles {
		if n == usrRole {
			roleIndex = s
		}
	}
	return roleIndex
}
