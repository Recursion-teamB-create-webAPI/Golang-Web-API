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
	JwtSecretKey   string
}

type ImageArray struct {
	Images [constants.SearchResultNumber]string `json:"images"`
}

type ResponseSearch struct {
	ImageData ImageArray `json:"imageData"`
	Status    string     `json:"status"`
	Cause     string     `json:"cause"`
}

type ResponseSignUp struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type ResponseSignIn struct {
	Username string `json:"username"`
	Token string `json:"token"`
	Status int `json:"status"`
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
	Id          int        `json:"id"`
	Item        string     `json:"item"`
	ImageData   ImageArray `json:"imageData"`
	SearchCount int        `json:"search_count"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

type User struct {
	Username string
	Password string
}
