package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "log"
)

type Item struct{
    ID      int    `json:"id"`
    Name    string `json:"namme"`
    URL     string `json:"url"`
}

var list = []Item{
    {ID: 1,Name: "Dog1",URL: "https://as2.ftcdn.net/v2/jpg/03/34/63/09/1000_F_334630968_IXzdadi2jSixyk9pNpQtyDo6XoNz1EnA.jpg"},
    {ID: 2,Name: "Dog2",URL: "https://t3.ftcdn.net/jpg/04/39/70/70/240_F_439707093_K3DVqaK8CiGv7XjAzKobujCw9DctQfZa.jpg"},
    {ID: 3,Name: "Dog3",URL: "https://t3.ftcdn.net/jpg/01/40/30/16/240_F_140301646_h3P1Nxiz3cfEtlAVXHwF45rpLg0Jh6tE.jpg"},
}




func listHandler(w http.ResponseWriter,r*http.Request){

    w.Header().Set("Content-TYpe","application/json")

    if err := json.NewEncoder(w).Encode(list); err != nil{
        log.Fatalf("Could not encode response: %v",err)
    }

}



func main(){
    fmt.Println("Starting the server!")
    log.Println("Starting server on port 8080...")
    
    // ルートとハンドラ関数を定義
    http.HandleFunc("/api/list",listHandler)

    // 8000番ポートでサーバを開始
    http.ListenAndServe(":8000", nil)
}
