package views

import (
	"database/sql"
	"encoding/json"

	"github.com/MaxRubel/kennel-server-go/models"
)

func GetAllLocations() ([]byte, error) {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM Location")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var locations []models.Location

	for rows.Next() {
		var location models.Location
		err := rows.Scan(&location.Id, &location.Name, &location.Address)
		if err != nil {
			panic(err)
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		panic(err)

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
