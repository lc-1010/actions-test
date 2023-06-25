package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
