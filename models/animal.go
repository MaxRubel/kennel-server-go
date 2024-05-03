package models

type Animal struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Breed      string `json:"breed"`
	LocationID int    `json:"location_id"`
	CustomerID int    `json:"customer_id"`
}
