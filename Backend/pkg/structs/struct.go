package structs

import "database/sql"

var db *sql.DB

type Item struct {
	ID int `json:"id"`
	Name string `json:"name`
	URL string `json:"url`
}

type Response struct{
	Status string `json:"id"`
	Items []Item `json:"items`
}
