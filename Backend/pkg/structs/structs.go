package structs

import (
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
)

type Env struct {
	SearchEngineId string
	CsePath        string
	PortNumber     string
}

type ResponseImage struct {
	Images [constants.SearchResultNumber]string `json:"images"`
}
