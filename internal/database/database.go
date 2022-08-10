package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Open() {
	cfg := mysql.Config{
		User:   "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB",
	}

	_db, err := sql.Open("mysql", cfg.FormatDSN())

	if err == nil {
		fmt.Printf("No errero")
	}

	Db = _db
}
