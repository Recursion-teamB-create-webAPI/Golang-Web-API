package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/handlers"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
)

func main() {
	// .envから値を取得する
	env := utils.GetEnvData()

	// MongoDBサーバーへのタイムアウト設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// MongoDBサーバーへの接続
	mydb := &dao.Database{}
	mydb.Connect(env, ctx)
	defer mydb.Disconnect(ctx)

	// 接続確認
	mydb.Ping(ctx)

	// データベースとコレクションの選択
	mydb.SetDbCol(env)

	// 初期データ投入
	mydb.InsertInitData()

	log.Println("Starting the server!")

	// 検索用エンドポイントにアクセスされたら呼び出す
	http.Handle("/api/search", http.HandlerFunc(handlers.SearchHandler(env, mydb)))

	// 8000番ポートでサーバを開始
	http.ListenAndServe(env.PortNumber, nil)
}
