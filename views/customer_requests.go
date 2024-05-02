package views

import (
	"database/sql"
	"encoding/json"

	"github.com/MaxRubel/kennel-server-go/models"
)

func GetAllCustomers() ([]byte, error) {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM Customer")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.Password)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		panic(err)

	}
	locationsJson, err := json.Marshal(customers)
	if err != nil {
		panic(err)

	}
	return locationsJson, nil
}

func GetSingleCustomer(id int) ([]byte, error) {
	var customer models.Customer

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.QueryRow("SELECT * FROM Customer WHERE Id = ?", id).
		Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.Password)

	if err != nil {
		panic(err)
	}

	customerJson, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	return customerJson, nil
}
