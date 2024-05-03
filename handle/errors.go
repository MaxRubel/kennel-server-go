package handle

import (
	"database/sql"
	"fmt"
	"time"
)

func Errors(msg string){

	db, err := sql.Open("sqlite3", "./errors.sqlite3")

	if err != nil {
		fmt.Println("Unable to open error log database")
		return
	}

	defer db.Close()

	currenTime := time.Now()

	query := `INSERT INTO Error (time, message) VALUES (?, ?);`

	_, err = db.Exec(query, currenTime, msg)

	if err != nil {
		fmt.Print("Error inserting new error message into error log:", err)
		return
	}

	fmt.Println("Error logged...")
}

func TestErrorDb(){
	db, err := sql.Open("sqlite3", "./errors.sqlite3")

	if err != nil {
		fmt.Println("Unable to open error log database")
		return
	}

	defer db.Close()
	var result string
	err = db.QueryRow("SELECT message FROM Error WHERE Id = 1").Scan(&result)
	if err != nil {
		fmt.Print("Unable to query error logs")
	}
	fmt.Println(result)
}