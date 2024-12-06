package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "topfijUf5!"
	dbname   = "testDb"
)

type Cat struct {
	Name string
	Age  int
}

func Init() {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	CheckError(err)

	fmt.Println("Successfully connected to PostgreSQL db.")
}

func Close() {
	err := db.Close()
	if err != nil {
		fmt.Println("Error closing DB")
	}
}

func GetCats() ([]Cat, error) {
	rows, err := db.Query("SELECT * FROM cats")
	CheckError(err)
	//fmt.Println(rows.Columns())

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println("Error closing rows")
		}
	}(rows)

	var cats []Cat

	for rows.Next() {
		var cat Cat
		CheckError(rows.Scan(&cat.Name, &cat.Age))
		cats = append(cats, cat)
	}

	if err = rows.Err(); err != nil {
		return cats, err
	}

	return cats, nil
}

func DbMain() {
	Init()
	defer fmt.Println("Connection closed!")
	defer Close()

	//_, err := db.Exec("INSERT INTO cats (name, age) VALUES ('Ark', 19)")
	//CheckError(err)
	result, err := GetCats()
	CheckError(err)
	fmt.Println(result)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
