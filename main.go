package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var username, password, name, host, port, tableName string

	fmt.Print("db_username:")
	fmt.Scan(&username)

	fmt.Print("db_password:")
	fmt.Scan(&password)

	fmt.Print("db_name:")
	fmt.Scan(&name)

	fmt.Print("db_host:")
	fmt.Scan(&host)

	fmt.Print("db_port:")
	fmt.Scan(&port)

	fmt.Print("db_table_name (table to display rows):")
	fmt.Scan(&tableName)
	fmt.Print("-----------------------------------")

	// Connect to the RDS database
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
}
