package db

import (
	"database/sql"
)

type Database interface {
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...any) (sql.Result, error)
	//	ExecRow(string, ...interface{}) (*sql.Result, error)
}

type Queries struct {
	db Database
}

func New(db Database) *Queries {
	return &Queries{db}
}
