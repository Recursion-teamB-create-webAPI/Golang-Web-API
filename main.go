package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func listHandler(w http.ResponseWriter,r*http.Request){
    query := r.URL.Query()
    keyword := query.Get("keyword")

    f,err := os.Open("sample.json")
    if err != nil{
        fmt.Println("error")
    }

    defer f.Close()

    response := map[string]string{
        "list": "test" + keyword,
    }

    w.Header().Set("Content-TYpe","application/json")

    json.NewEncoder(w).Encode(response)
}



func main(){
    fmt.Println("Starting the server!")
    
    // ルートとハンドラ関数を定義
    http.HandleFunc("/api/list",listHandler)

    // 8000番ポートでサーバを開始
    http.ListenAndServe(":8000", nil)
}
