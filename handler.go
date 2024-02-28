package handler

import (
	"encoding/json"
	"net/http"
)

type List struct{
	ID int
	Name string
	imageURL string
}


func listHandler(w http.ResponseWriter,r*http.Request){
    query := r.URL.Query()
    keyword := query.Get("keyword")

    response := map[string]string{
        "list": "test" + keyword,
    }

    w.Header().Set("Content-TYpe","application/json")

    json.NewEncoder(w).Encode(response)
}