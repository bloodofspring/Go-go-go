package databases

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	var err error

	connStr := ""
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PostgreSQL db.")
}
