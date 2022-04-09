package controller

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tubes_pbp")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
