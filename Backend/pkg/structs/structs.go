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

type ResponseList struct {
	List   []string `json:"list"`
	Status string   `json:"status"`
	Cause  string   `json:"cause"`
}

type ResponseTotalResult struct {
	TotalResult []TotalResultItems `json:"totalResult"`
	Status      string             `json:"status"`
	Cause       string             `json:"cause"`
}

type TotalResultQueryArray struct {
	Page    int
	PerPage int
	Order   string
}

type ImageArray struct {
	Images [constants.SearchResultNumber]string `json:"images"`
}

type InitImageItems struct {
	ImageItems []Items `json:"ImageItems"`
}

type TotalResultItems struct {
	Item        string `json:"item"`
	SearchCount int    `json:"search_count"`
	UpdatedAt   string `json:"updated_at"`
}

type Items struct {
	Item      string     `json:"item"`
	ImageData ImageArray `json:"imageData"`
}

type DatabaseImage struct {
	Id          int        `json:"id"`
	Item        string     `json:"item"`
	ImageData   ImageArray `json:"imageData"`
	SearchCount int        `json:"search_count"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}
