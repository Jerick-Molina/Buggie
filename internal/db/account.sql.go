package database

import "database/sql"

type createAccByCompanyCodeParams struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	CompanyId int    `json:"companyId"`
}

const createAccByCompanyCode = `-- Creates Account based on the Company's invite code
insert into Users(
	FirstName,
	LastName,
	Email,
	Password,
	Role,
	CompanyId
) values(?,?,?,?,?,?) returning UserId,Role,CompanyId
`

func (db *Queries) createAccountByCompanyCode(arg createAccByCompanyCodeParams) (Account, error) {
	row := db.actions.QueryRow(createAccByCompanyCode, arg.FirstName, arg.LastName, arg.Email, arg.Password, arg.Role, arg.CompanyId)
	var u Account
	err := row.Scan(
		&u.Id,
		&u.Role,
		&u.CompanyId,
	)
	return u, err
}

type searchAccsByCompanyIdParams int

const searchAccsByCompanyId = `-- Finds all users by associated Company Id

`

func (db *Queries) searchAccountsByCompanyId(arg searchAccsByCompanyIdParams) ([]Account, error) {
	var usrs []Account
	rows, err := db.actions.Query(searchAccsByCompanyId, arg)
	for rows.Next() {
		var u Account
		if err := rows.Scan(
			&u.FirstName,
			&u.LastName,
			&u.Role,
		); err != nil {
			return nil, err
		}

		usrs = append(usrs, u)
	}
	return usrs, err
}

type searchEmailExistValidationParam string

const searchAccountValidation = `-- Checks to see if Email exist if so user cannot create a account with that email
select Email from Users where Email = ?
`

func (db *Queries) searchEmailExistValidation(arg searchEmailExistValidationParam) (bool, error) {

	row := db.actions.QueryRow(searchAccountValidation, arg)
	var u Account
	err := row.Scan(&u.Email)

	if err == sql.ErrNoRows {
		return true, err
	}

	return false, err
}
