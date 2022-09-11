package main

import (
	"database/sql"

	"Buggie_2.0/cmd/server"
	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{User: "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB"}
	dbs, err := sql.Open("mysql", cfg.FormatDSN())

	db := database.NewStore(dbs)
	if err != nil {

	}

	if err := server.Run(db); err != nil {

	}

}
