package main

import (
	"database/sql"
	"log"
	"net/http"

    "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/handlers"
)

func main() {
	// MySQLに接続
	var err error
	db, err = sql.Open("mysql", "ユーザ名:パスワード@tcp(ホスト:ポート)/データベース名")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ルーターの設定
	r := mux.NewRouter()
	r.HandleFunc("/api/list",listHandler).Methods("GET")

	// サーバーの起動
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}


