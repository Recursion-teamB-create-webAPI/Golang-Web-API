package structs

import (
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
)

type Env struct {
	SearchEngineId string
	KeyFileName    string
	CsePath        string
	PortNumber     string
	DatabaseName   string
	MysqlUri       string
}

type ImageArray struct {
	Images [constants.SearchResultNumber]string `json:"images"`
}

type ResponseSearch struct {
	ImageData ImageArray `json:"imageData"`
	Status    string     `json:"status"`
	Cause     string     `json:"cause"`
}

type ResponseDescription struct {
	Description DatabaseImage `json:"description"`
	Status      string        `json:"status"`
	Cause       string        `json:"cause"`
}

type InitImageItems struct {
	ImageItems []Items `json:"ImageItems"`
}

type Items struct {
	Item      string     `json:"item"`
	ImageData ImageArray `json:"imageData"`
}

type DatabaseImage struct {
	Id          string                               `json:"id"`
	Item        string                               `json:"item"`
	Images      [constants.SearchResultNumber]string `json:"images"`
	SearchCount int                                  `json:"search_count"`
	CreatedAt   string                               `json:"created_at"`
	UpdatedAt   string                               `json:"updated_at"`
}
