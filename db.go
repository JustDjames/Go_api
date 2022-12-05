package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func check_conn(db *sql.DB) error {
	// use Ping to check that DB connection still exists and establish a connection if it doesn't

	debug.Println("checking for Database connections")
	e := db.Ping()

	return e
}

func check_table(db *sql.DB, table string) error {
	e := check_conn(db)

	if e != nil {
		return e
	}

	// check for table

	// check columns in table are correct datatypes
}
