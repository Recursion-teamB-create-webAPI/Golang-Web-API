package main

import (
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/handlers"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
)

func main() {
	// .envから値を取得する
	env := utils.GetEnvData()

	log.Println("Starting the server!")

	// 検索用エンドポイントにアクセスされたら呼び出す
	http.Handle("/api/search", http.HandlerFunc(handlers.SearchHandler(env)))

	// 8000番ポートでサーバを開始
	http.ListenAndServe(env.PortNumber, nil)
}
