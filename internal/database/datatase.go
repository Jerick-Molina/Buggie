package databasess

// import (
// 	"database/sql"
// 	"fmt"
// )

// type DatabaseServer struct {
// 	*sql.DB
// }

// func Run(db *sql.DB) *DatabaseServer {
// 	return &DatabaseServer{db}
// }

// // func Run() error {
// // 	cfg := mysql.Config{User: "buggie",
// // 		Passwd: "Mixon9090",
// // 		Net:    "tcp",
// // 		Addr:   "192.168.3.139",
// // 		DBName: "BuggieDB"}
// // 	db, err := sql.Open("mysql", cfg.FormatDSN())

// // 	if err != nil {
// // 		return nil
// // 	}

// // 	return nil
// // }
// func (db *DatabaseServer) TestDB(usrname string) {
// 	sqlQuierie := "select FirstName from Users where Email = ?"
// 	var e string
// 	if sql := db.QueryRow(sqlQuierie, usrname).Scan(&e); sql != nil {
// 		fmt.Println(sql.Error())
// 	}

// 	fmt.Println(e)

// }
