package views

import (
	"database/sql"
	"encoding/json"

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
		err := rows.Scan(&animal.Id, &animal.Name, &animal.Breed, &animal.Status, &animal.LocationID, &animal.CustomerID)
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
		Scan(&animal.Id, &animal.Name, &animal.Breed, &animal.Status, &animal.LocationID, &animal.CustomerID)

	if err != nil {
		panic(err)
	}

	animalJson, err := json.Marshal(animal)
	if err != nil {
		panic(err)
	}
	return animalJson, nil
}
