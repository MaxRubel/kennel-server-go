package models

type Animal struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Breed      string `json:"breed"`
	Status     string `json:"status"`
	LocationID int    `json:"location_id"`
	CustomerID int    `json:"customer_id"`
}
