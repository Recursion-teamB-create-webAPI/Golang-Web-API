package structs

import (
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
)

type Env struct {
	SearchEngineId  string
	CsePath         string
	PortNumber      string
	MongoUri        string
	MongoDatabase   string
	MongoCollection string
}

type ImageArray struct {
	Images [constants.SearchResultNumber]string `json:"images"`
}

type ResponseImage struct {
	ImageData ImageArray `json:"imageData"`
	Status    string     `json:"status"`
}

type InitImageItems struct {
	ImageItems []Items `json:"items"`
}

type Items struct {
	Item      string     `json:"item"`
	ImageData ImageArray `json:"imageData"`
}

type DatabaseImage struct {
	Id          int       `bson:"id"`
	Item        string    `bson:"item"`
	Images      []string  `bson:"images"`
	SearchCount int       `bson:"search_count"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
