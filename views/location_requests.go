package views

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/MaxRubel/kennel-server-go/handle"
	"github.com/MaxRubel/kennel-server-go/models"
)

func getAnimalsOfLocation(location_id int)([]models.Animal, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		handle.Errors("Unable to open SQL DB get animals of location")
		return nil, errors.New("nope")
	}
	rows, err := db.Query("SELECT * FROM Animal WHERE location_id = ?", location_id)
	if err != nil {
		handle.Errors("Unable to query db get all animals of location")
		return nil, errors.New("nope")
	}
	var animals []models.Animal
	for rows.Next(){
		var animal models.Animal
		rows.Scan(&animal.Id, &animal.Name, &animal.Breed, &animal.Status, &animal.LocationID, &animal.CustomerID)
		animals = append(animals, animal)
	}
	return animals, nil
}

func getEmployeesOfLocation(location_id int)([]models.Employee, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		handle.Errors("Unable to open SQL DB get employees of location")
		return nil, errors.New("nope")
	}
	rows, err := db.Query("SELECT * FROM Employee WHERE location_id = ?", location_id)
	if err != nil {
		handle.Errors("Unable to query db get all animals of location")
		return nil, errors.New("nope")
	}
	var employees []models.Employee
	for rows.Next(){
		var employee models.Employee
		rows.Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Location_id)
		employees = append(employees, employee)
	}
	return employees, nil
}

func GetAllLocations() ([]byte, error) {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		handle.Errors("Unable to open SQL DB get all locations")
		return nil, errors.New("unable to open sql db to get all locations")
	}

	rows, err := db.Query("SELECT * FROM Location")
	if err != nil {
		handle.Errors("Unable to query db get all locations")
		return nil, errors.New("unable to query db get all locations")
	}

	defer rows.Close()

	var locations []models.Location

	for rows.Next() {
		var location models.Location
		err := rows.Scan(&location.Id, &location.Name, &location.Address)
		if err != nil {
			handle.Errors("Unable to scan query result get all locations")
			return nil, errors.New("unable to scan query result get all locations")
		}
		locations = append(locations, location)
	}

	for i := range locations {
		location := locations[i] 
		animals, err := getAnimalsOfLocation(location.Id)
		if err != nil {
			handle.Errors("Error getting animals for location with ID")
			continue
		}
		location.Animals = animals
		locations[i] = location
	}

	for i := range locations {
		location := locations[i] 
		employees, err := getEmployeesOfLocation(location.Id)
		if err != nil {
			handle.Errors("Error getting animals for location with ID")
			continue
		}
		location.Employees = employees
		locations[i] = location 
	}

	locationsJson, err := json.Marshal(locations)
	if err != nil {
		panic(err)
	}
	return locationsJson, nil
}

func GetSingleLocation(id int) ([]byte, error) {
	var location models.Location

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.QueryRow("SELECT * FROM Location WHERE Id = ?", id).
		Scan(&location.Id, &location.Name, &location.Address)

	if err != nil {
		panic(err)
	}

	locationJson, err := json.Marshal(location)
	if err != nil {
		panic(err)
	}
	return locationJson, nil
}
