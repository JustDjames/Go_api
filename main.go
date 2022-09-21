package main

import (
	"flag"
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
	api_port_arg := flag.String("api_port", "", "The port the api will listening on")

	flag.Parse()

	db_hostname := var_parser(*db_hostname_arg, "DB_HOSTNAME", "")
	db_pass := var_parser(*db_pass_arg, "DB_PASS", "")
	db_user := var_parser(*db_user_arg, "DB_USER", "")
	db_port := var_parser(*db_port_arg, "DB_PORT", "")
	db_name := var_parser(*db_name_arg, "DB_NAME", "")
	api_port := var_parser(*api_port_arg, "API_PORT", "8080")

	info.Println("db_hostname:", db_hostname)
	info.Println("db_pass:", db_pass)
	info.Println("db_user:", db_user)
	info.Println("db_port:", db_port)
	info.Println("db_name:", db_name)
	info.Println("dapi_port", api_port)
}
