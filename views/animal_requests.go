package views

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

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
			panic(err)
		}
		animals = append(animals, animal)
	}

	if err := rows.Err(); err != nil {
		panic(err)

	}
	animalsJson, err := json.Marshal(animals)
	if err != nil {
		panic(err)

	}
	return animalsJson, nil

}

func GetSingleAnimal(id int) ([]byte, error) {
	var animal models.Animal

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.QueryRow("SELECT * FROM Animal WHERE Id = ?", id).
		Scan(&animal.Id, &animal.Name, &animal.Status, &animal.Breed, &animal.LocationID, &animal.CustomerID)

	if err != nil {
		panic(err)
	}

	animalJson, err := json.Marshal(animal)
	if err != nil {
		panic(err)
	}
	return animalJson, nil
}

func CreateNewAnimal(newAnimal models.Animal) error {
	db, err := sql.Open("sqlite3", "./db.sqlite3")

	if err != nil {
		panic(err)
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
		return errors.New("error inserting new animal into DB")
	} else {
		fmt.Print("Animal sucessfully created")
		return nil
	}
}
