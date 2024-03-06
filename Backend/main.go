package main

import (
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/handlers"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// .envから値を取得する
	env := utils.GetEnvData(constants.BeforeLevel1)

	// MySQLに接続する
	mydb := &dao.Database{}
	mydb.Connect(env)
	defer mydb.UseDb.Close()

	// 接続確認
	err := mydb.UseDb.Ping()
	if err != nil {
		log.Println("Database connection failed")
	} else {
		log.Println("Database connection successful")
	}

	// テーブルの作成
	mydb.CreateTable()

	// 初期データ投入
	mydb.InsertInitData(constants.BeforeLevel0)

	log.Println("Starting the server!")

	// 各エンドポイントにアクセスされたら呼び出す
	http.Handle("/api/search", http.HandlerFunc(handlers.SearchHandler(env, mydb)))

	http.Handle("/api/description", http.HandlerFunc(handlers.DescriptionHandler(mydb)))

	http.Handle("/api/signup", http.HandlerFunc(handlers.SignUpHandler(env, mydb)))

	http.Handle("/api/signin", http.HandlerFunc(handlers.SignInHandler(env, mydb)))

	// 8000番ポートでサーバを開始
	http.ListenAndServe(env.PortNumber, nil)

}
