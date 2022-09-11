package database

import (
	"database/sql"
)

type Databases interface {
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
	//	ExecRow(string, ...interface{}) (*sql.Result, error)
}

type Queries struct {
	actions Databases
}
type Store struct {
	*Queries
	db *sql.DB
}

func New(db Databases) *Queries {
	return &Queries{actions: db}
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
