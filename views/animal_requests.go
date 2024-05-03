package views

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/MaxRubel/kennel-server-go/handle"
	"github.com/MaxRubel/kennel-server-go/models"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllAnimals() ([]byte, error) {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM Animal")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var animals []models.Animal

	for rows.Next() {
		var animal models.Animal
		err := rows.Scan(&animal.Id, &animal.Name, &animal.Status, &animal.Breed, &animal.LocationID, &animal.CustomerID)
		if err != nil {
			handle.Errors("error scanning animal from DB query")
			return nil, errors.New("error scanning animal from db query")
		}
		animals = append(animals, animal)
	}

	animalsJson, err := json.Marshal(animals)
	if err != nil {
		handle.Errors("error marshalling json get all animals")
		return nil, errors.New("error marshalling json")

	}
	return animalsJson, nil

}

func GetSingleAnimal(id int) ([]byte, error) {
	var animal models.Animal

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		handle.Errors("Error opening DB")
		return nil, errors.New("error opening db")
	}

	defer db.Close()
	err = db.QueryRow("SELECT * FROM Animal WHERE Id = ?", id).
		Scan(&animal.Id, &animal.Name, &animal.Status, &animal.Breed, &animal.LocationID, &animal.CustomerID)

	if err != nil {
		handle.Errors("Getting from DB table: Animal")
		return nil, errors.New("error getting from DB table: animal")
	}

	animalJson, err := json.Marshal(animal)
	if err != nil {
		handle.Errors("Error marshalling JSON from Animal")
		return nil, errors.New("error getting from DB table: animal")
	}
	return animalJson, nil
}

func CreateNewAnimal(newAnimal models.Animal) error {
	db, err := sql.Open("sqlite3", "./db.sqlite3")

	if err != nil {
		handle.Errors("Error opening DB create new animal")
		return errors.New("error opening db create new animal")
	}

	defer db.Close()

	query := `
	INSERT INTO Animal
		(name, breed, status, location_id, customer_id)
	VALUES
		($1, $2, $3, $4, $5)
`
	_, err = db.Exec(query, newAnimal.Name, newAnimal.Status, newAnimal.Breed, newAnimal.LocationID, newAnimal.CustomerID)
	if err != nil {
		handle.Errors("error inserting new animal into DB")
		return errors.New("error inserting new animal into DB")
	} else {
		return nil
	}
}
