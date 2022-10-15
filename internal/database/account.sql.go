package db

import (
	"database/sql"

	"github.com/Jerick-Molina/Buggie/internal/security"
)

type CreateAccByCompanyCodeParams struct {
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
) values(?,?,?,?,?,?);
`

func (q *Queries) CreateAccountByCompany(arg CreateAccByCompanyCodeParams) (Account, error) {
	arg.Role = "Admin"
	arg.Password = security.HashPassword(arg.Password)
	_, err := q.db.Exec(createAccByCompanyCode, arg.FirstName, arg.LastName, arg.Email, arg.Password, arg.Role, arg.CompanyId)
	if err != nil {
		return Account{}, err
	}
	var u Account

	if result := q.db.QueryRow("select CompanyId,UserId,Role from Users where Email = ?", arg.Email).Scan(&u.CompanyId, &u.Id, &u.Role); result != nil {
		return Account{}, result
	}
	return u, err
}

func (q *Queries) CreateAccountByCompanyCode(arg CreateAccByCompanyCodeParams) (Account, error) {
	row := q.db.QueryRow(createAccByCompanyCode, arg.FirstName, arg.LastName, arg.Email, arg.Password, arg.Role, arg.CompanyId)
	var u Account

	err := row.Scan(
		&u.Id,
		&u.Role,
		&u.CompanyId,
	)
	return u, err
}

type SearchAccsByCompanyIdParams int

const searchAccsByCompanyId = `-- Finds all users by associated Company Id
`

func (q *Queries) SearchAccountsByCompanyId(arg SearchAccsByCompanyIdParams) ([]Account, error) {
	var usrs []Account
	rows, err := q.db.Query(searchAccsByCompanyId, arg)
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

const searchAccountValidation = `-- Checks to see if Email exist if so user cannot create a account with that email
select Email from Users where Email = ?
`

func (q *Queries) SearchEmailExistValidation(email string) error {

	row := q.db.QueryRow(searchAccountValidation, email)
	var u Account
	err := row.Scan(&u.Email)

	if err != sql.ErrNoRows {
		return sql.ErrNoRows
	}

	return nil
}

const validSignInCredentials = `--Checks to see if credidentials is valid
select UserId,Role,CompanyId from Users where Email = ? and Password = ?
`

func (q *Queries) ValidSignInCredentials(email string, password string) (Account, error) {
	var acc Account
	if result := q.db.QueryRow(validSignInCredentials, email, password).Scan(&acc.Id, &acc.Role, &acc.CompanyId); result != nil {
		//Credidentials invalid
		return Account{}, sql.ErrNoRows
	}

	return acc, nil
}
