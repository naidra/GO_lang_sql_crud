package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	// cfg := mysql.Config{
	// 	User:   "sql7610344", // os.Getenv("DBUSER")
	// 	Passwd: "zhwh4d66Xe", // os.Getenv("DBPASS")
	// 	Net:    "tcp",
	// 	Addr:   "sql7.freesqldatabase.com:3306",
	// 	DBName: "sql7610344",
	// }
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", "golang_user:golang_user@tcp(localhost:3306)/mysql")

	// user := "sql7610344"
	// pass := "zhwh4d66Xe"
	// host := "sql7.freesqldatabase.com:3306"
	// dbname := "sql7610344"
	// connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	// db, err := sql.Open("mysql", connection)
	// if err != nil {
	// 	panic(err)
	// }

	db, err := sql.Open("mysql", "sql7610344:zhwh4d66Xe@tcp(sql7.freesqldatabase.com:3306)/sql7610344")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
