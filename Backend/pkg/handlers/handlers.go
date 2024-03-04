package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func SearchHandler(env structs.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // 許可するオリジンを指定
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   // 許可するHTTPメソッドを指定
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // 許可するリクエストヘッダーを指定

		fmt.Println("Search Handler")

		query := r.URL.Query()
		keyword := query.Get("keyword")
		fmt.Println("keyword>>", keyword)
		jsonData, err := os.ReadFile("search-key.json")
		if err != nil {
			log.Println(err.Error)
		}

		conf, err := google.JWTConfigFromJSON(jsonData, env.CsePath)
		if err != nil {
			log.Println(err)
		}

		client := conf.Client(context.Background())
		cseService, err := customsearch.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			return
		}
		// 検索ワードの設定
		search := cseService.Cse.List().Q(keyword)

		// 検索エンジンIDを設定
		search.Cx(env.SearchEngineId)
		// Custom Search Engineで「画像検索」をオンにする
		search.SearchType("image")

		search.Start(1)
		call, err := search.Do()
		if err != nil {
			log.Println(err)
		}

		var response structs.ResponseImage

		for index, r := range call.Items {
			response.Images[index] = r.Link
		}

		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}
