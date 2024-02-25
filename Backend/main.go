package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを解析する
	query := r.URL.Query()
	name := query.Get("name") // "名前"から"name"へ修正

	// レスポンス用のマップを作成
	response := map[string]string{
		"message": "Hello " + name, // “message”： "Hello " + name、から修正
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを解析する
	query := r.URL.Query()
	keyword := query.Get("keyword")

	// レスポンス用のマップを作成
	response := map[string]string{}

	jsonData, err := os.ReadFile("search-key.json")
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(jsonData, "https://www.googleapis.com/auth/cse")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(context.Background())
	cseService, err := customsearch.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}
	// 検索ワードの設定
	search := cseService.Cse.List().Q(keyword)

	// 検索エンジンIDを設定
	search.Cx(getEnvData())
	// Custom Search Engineで「画像検索」をオンにする
	search.SearchType("image")

	search.Start(1)
	call, err := search.Do()
	if err != nil {
		log.Fatal(err)
	}

	for index, r := range call.Items {
		response[strconv.Itoa(index+1)] = r.Link
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func getEnvData() string {
	// envファイルのパスを渡す
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// .envから値を取得する
	searchEnginId := os.Getenv("SEARCH_ENGINE_ID")

	return searchEnginId
}

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/hello", helloHandler)

	// 検索用エンドポイントにアクセスされたら呼び出す
	http.HandleFunc("/api/search", searchHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
