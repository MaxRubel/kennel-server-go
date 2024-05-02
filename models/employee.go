package models

type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Location_id int    `json:"location_id"`
}
