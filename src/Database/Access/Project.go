package access

import (
	"fmt"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	database "github.com/Jerick-Molina/Buggie/src/Database"
)

func ProjectSearchByCompanyId(companyId int) {

	sql := "select * from Users"
	var test []properties.User
	db := database.Access()
	if db == nil {

	}

	rows, errs := db.Query(sql)

	if errs != nil {
		fmt.Println(errs.Error())
	}
	e, errs := rows.Columns()

	fmt.Println(e)
	fmt.Println(rows.Next())

	defer rows.Close()

	for rows.Next() {

		var usrs properties.User

		err := rows.Scan(&usrs.Id, &usrs.FirstName, &usrs.LastName, &usrs.Email, &usrs.Password, &usrs.Role, &usrs.CompanyId)
		if err != nil {
			fmt.Println(err.Error())
		}
		test = append(test, usrs)
	}

	fmt.Println(test)
}

func CreateProject() {
	sql := ""

	fmt.Println(sql)
}

func EditProject() {
	sql := ""

	fmt.Println(sql)
}

func DeleteProject() {
	sql := ""

	fmt.Println(sql)
}
