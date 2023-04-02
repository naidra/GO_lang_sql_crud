package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var con *sql.DB

func Connect() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "sql7610344", // os.Getenv("DBUSER")
		Passwd: "zhwh4d66Xe", // os.Getenv("DBPASS")
		Net:    "tcp",
		Addr:   "sql7.freesqldatabase.com:3306",
		DBName: "sql7610344",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", "golang_user:golang_user@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
