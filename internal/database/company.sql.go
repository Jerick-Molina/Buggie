package db

import (
	"database/sql"
	"fmt"

	"github.com/Jerick-Molina/Buggie/internal/database/helper"
)

const createCompany = `insert into Company(Name,CompanyCode) values (?,?)
`

func (q *Queries) CreateCompany(companyName string) (int, error) {
	codeExist := true
	var code string
	var compId int

	//Creates code and checks if code exist
	for codeExist == true {

		code = helper.CreateCompanyCode()

		if results := q.db.QueryRow(helper.CodeSearch, code).Scan(); results == sql.ErrNoRows {
			codeExist = false

		}

	}
	_, err := q.db.Exec(createCompany, companyName, code)
	if err != nil {
		//Errr
		fmt.Println("hee")
		return -1, err
	}

	if r := q.db.QueryRow("select CompanyId from Company where CompanyCode = ?", code).Scan(&compId); r != nil {

	}
	fmt.Println(compId)
	return compId, nil
}

const findCompanyByCode = `--finds company by company code
select CompanyCode from Company where CompanyCode = ?
`

func (q *Queries) FindCompanyByCode(code string) (int, error) {

	if err := q.db.QueryRow(findCompanyByCode, code).Scan(); err != sql.ErrNoRows {

		return -1, sql.ErrNoRows
	}

	return -1, nil
}
