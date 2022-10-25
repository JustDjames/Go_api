package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// global var that allows application access to Db handler
var db *sql.DB

func db_init(user string, pass string, hostname string, port string, db_name string) error {

	// creating the connection string from the arguments

	debug.Print("creating database handler")

	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, hostname, port, db_name)

	db, e := sql.Open("mysql", connection_string)

	if e != nil {
		err.Printf("error encountered when creating handle for database: %s", e)
		os.Exit(1)
	}

	// check that connection to database is still alive

	// debug.Print("checking database connection")

	// e = db.Ping()

	// if e != nil {

	// }

	//  returning Ping which will check the connection to the database
	return db.Ping()
}
