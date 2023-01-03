package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// creating custom Loggers
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger

	// creating db vars
	db_hostname string
	db_pass     string
	db_user     string
	db_port     string
	db_name     string
	db_table    string
	api_port    string
)

func init() {
	// configuring custom loggers
	debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	err = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func var_parser(arg string, env_var string, default_var string) (v string) {

	if arg != "" {
		v = arg
	} else {
		//  check env var exists if arg is empty
		env, exist := os.LookupEnv(env_var)

		if exist {
			v = env
		} else if default_var != "" {
			info.Printf("Environment Variable %s does not exist. Using default value %s ", env_var, default_var)

			v = default_var
		} else {
			err.Printf("%s not set in either environment vars or arguments and no default value availble.\n\n", env_var)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
	return v
}

func main() {
	log.Println("parsing environment variables and arguments")

	db_hostname_arg := flag.String("db_hostname", "", "The hostname of the database")
	db_pass_arg := flag.String("db_pass", "", "The password for the user used connect to database")
	db_user_arg := flag.String("db_user", "", "The user used to connect to database")
	db_port_arg := flag.String("db_port", "", "The database port")
	db_name_arg := flag.String("db_name", "", "The name of the database")
	db_table_arg := flag.String("db_table", "", "The name of the table used in the database. defaults to same value as the db_name argument")
	api_port_arg := flag.String("api_port", "", "The port the api will listening on. defaults to 8080")

	flag.Parse()

	db_hostname := var_parser(*db_hostname_arg, "DB_HOSTNAME", "")
	db_pass := var_parser(*db_pass_arg, "DB_PASS", "")
	db_user := var_parser(*db_user_arg, "DB_USER", "")
	db_port := var_parser(*db_port_arg, "DB_PORT", "")
	db_name := var_parser(*db_name_arg, "DB_NAME", "")
	db_table := var_parser(*db_table_arg, "DB_TABLE", db_name)
	api_port := var_parser(*api_port_arg, "API_PORT", "8080")

	debug.Println("db_hostname:", db_hostname)
	debug.Println("db_pass:", db_pass)
	debug.Println("db_user:", db_user)
	debug.Println("db_port:", db_port)
	debug.Println("db_name:", db_name)
	debug.Println("db_table:", db_table)
	debug.Println("api_port", api_port)

	// creating the connection string from the arguments

	debug.Print("creating database handler")

	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass, db_hostname, db_port, db_name)

	db, e := sql.Open("mysql", connection_string)

	if e != nil {
		err.Printf("error encountered when creating handle for database: %s", e)
		os.Exit(1)
	}

	e = check_conn(db)

	if e != nil {
		err.Printf("error encountered when checking database connection: %s", e)
		os.Exit(1)
	} else {
		info.Print("Database connection created!")
	}

	e, table_exist, columns_correct := check_table(db, db_table)

	if e != nil {
		err.Printf("error encountered when checking %s table in database: %s", db_table, e)
		os.Exit(1)
	} else if table_exist == false {
		err.Printf("table %s doesn't exists in database!", db_table)
		os.Exit(1)
	} else if columns_correct == false {
		err.Printf("Table %s colums are incorrectly configured!", db_table)
		os.Exit(1)
	} else {
		info.Printf("table %s exists in database!", db_table)
	}
}
