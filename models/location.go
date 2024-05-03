package models

type Location struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Animals []Animal `json:"animals"`
	Employees []Employee  `json:"employees"`
}
