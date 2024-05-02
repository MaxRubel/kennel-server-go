package views

import (
	"database/sql"
	"encoding/json"

	"github.com/MaxRubel/kennel-server-go/models"
)

func GetAllEmployees() ([]byte, error) {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Location_id)
		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		panic(err)

	}
	employeesJson, err := json.Marshal(employees)
	if err != nil {
		panic(err)

	}
	return employeesJson, nil
}

func GetSingleEmployee(id int) ([]byte, error) {
	var employee models.Employee

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.QueryRow("SELECT * FROM Customer WHERE Id = ?", id).
		Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Location_id)

	if err != nil {
		panic(err)
	}

	employeeJson, err := json.Marshal(employee)
	if err != nil {
		panic(err)
	}
	return employeeJson, nil
}
