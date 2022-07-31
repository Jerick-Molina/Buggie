package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Open() {
	cfg := mysql.Config{
		User:   "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB",
	}

	_db, err := sql.Open("mysql", cfg.FormatDSN())
	errr := _db.Ping()
	if errr == nil {
		fmt.Printf("No errero")
	}

	if err != nil || errr != nil {
		fmt.Printf("No errero")
	}
	db = _db
}
func Access() *sql.DB {

	//first check if db is alive
	err := db.Ping()

	if err != nil {
		fmt.Println("DATABASE IS NOT CONNECTED")
		return nil
	}

	return db

}
