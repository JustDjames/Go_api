package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
structs i've created in order to contain the results from Database queries.

They use struct tags which seem to be used to carry meta-data for each of the struct fields.

I'm using it to add infor on how each struct field should be encoded/decoded to json format.

for more infor, look at: https://stackoverflow.com/a/30889373
*/
type Table struct {
	Field string `json:"Field"`
	Type  string `json:"Type"`
	Null  string `json:"Null"`
	Key   string `json:"Key"`
	// setting Default datatype as Nullstring as the value may or may not be null. it is a struct that contains a string which has the value if it exists, and a bool which is true if the string is not null
	Default sql.NullString `json:"Default"`
	Extra   string         `json:"Extra"`
}

func check_conn(db *sql.DB) error {
	// use Ping to check that DB connection still exists and establish a connection if it doesn't

	debug.Println("checking for Database connections")
	e := db.Ping()

	return e
}

func check_table(db *sql.DB, db_table string) (error, bool, bool) {
	e := check_conn(db)

	if e != nil {
		return e, false, false
	}

	// check for table

	result, e := db.Query("show tables;")

	if e != nil {
		return e, false, false
	}

	tables := make([]string, 0)
	for result.Next() {
		var table string

		// for each row, scan the result into the table var
		e := result.Scan(&table)
		if e != nil {
			return e, false, false
		}
		// append value in table to tables string slice
		tables = append(tables, table)
	}

	debug.Print(tables)

	table_exist := false
	for _, t := range tables {

		if db_table == t {
			table_exist = true
			break
		}
	}

	// check if table_exist is still false. if it is we don't need to check the fields (columns) as the table doesn't exist

	if table_exist == false {
		return e, table_exist, false
	}

	// creating slice with expected table field values

	expectedFields := make([]Table, 0)

	ID := Table{"ID", "varchar(225)", "NO", "PRI", sql.NullString{Valid: false}, ""}

	LastName := Table{"LastName", "varchar(225)", "NO", "", sql.NullString{Valid: false}, ""}

	FirstName := Table{"FirstName", "varchar(225)", "NO", "", sql.NullString{Valid: false}, ""}

	Occupation := Table{"Occupation", "varchar(225)", "NO", "", sql.NullString{Valid: false}, ""}

	DOB := Table{"DOB", "date", "NO", "", sql.NullString{Valid: false}, ""}

	expectedFields = append(expectedFields, ID, LastName, FirstName, Occupation, DOB)

	debug.Print(expectedFields)
	// check fields in table are correct datatypes

	result, e = db.Query(fmt.Sprintf("describe %s;", db_table))

	if e != nil {
		return e, table_exist, false
	}

	tableFields := make([]Table, 0)
	var table Table

	for result.Next() {
		e := result.Scan(&table.Field, &table.Type, &table.Null, &table.Key, &table.Default, &table.Extra)

		if e != nil {
			return e, table_exist, false
		}
		info.Print(table.Default.String)
		tableFields = append(tableFields, table)
	}
	debug.Print(tableFields)

	// compare tableFields and expectedFields

	columnsCorrect := false
	if len(expectedFields) == len(tableFields) {

		for i, field := range tableFields {

			if field == expectedFields[i] {
				columnsCorrect = true

				debug.Printf("%s configured as expected", field.Field)
			} else {
				err.Printf("%s not configured as expected", field.Field)
				columnsCorrect = false
				break
			}
		}

	}

	return e, table_exist, columnsCorrect
}
