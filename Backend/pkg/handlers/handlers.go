package handlers

import (
	"context"
	"encoding/json"
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
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		jsonData, err := os.ReadFile("search-key.json")
		if err != nil {
			log.Println(err)
		}

		conf, err := google.JWTConfigFromJSON(jsonData, env.CsePath)
		if err != nil {
			log.Println(err)
		}

		client := conf.Client(context.Background())
		cseService, err := customsearch.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Println(err)
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
