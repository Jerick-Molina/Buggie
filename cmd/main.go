package main

import (
	"database/sql"

	"github.com/Jerick-Molina/Buggie/cmd/server"
	db "github.com/Jerick-Molina/Buggie/internal/database"
	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User:   "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB"}
	dbConn, err := sql.Open("mysql", cfg.FormatDSN())

	store := db.NewStore(dbConn)
	if err != nil {

	}
	server := server.NewServer(store)

	server.Start("localhost:8080")
}
