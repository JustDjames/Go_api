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

func check_table(db *sql.DB, db_table string) (error, bool) {
	e := check_conn(db)

	if e != nil {
		return e, false
	}

	// check for table

	result, e := db.Query("show tables;")

	if e != nil {
		return e, false
	}

	tables := make([]string, 0)
	for result.Next() {
		var table string

		// for each row, scan the result into the table var
		e := result.Scan(&table)
		if e != nil {
			return e, false
		}
		// append value in table to tables string slice
		tables = append(tables, table)
	}

	debug.Print(tables)

	table_exist := false
	for _, t := range tables {

		if db_table == t {
			table_exist = true
		}
	}

	// if table exists, check columns in table are correct datatypes
	return e, table_exist
}
