package views

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func TestRequest() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var message string
		err := rows.Scan(&message)
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
}
