package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func SearchHandler(env structs.Env, mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseImage
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			log.Println("Data could not be retrieved because query parameters were not set")
			response.Status = "failed"
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				mydb.Update(keyword)
				response.ImageData.Images = [constants.SearchResultNumber]string(img.Images)
			} else {
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

				for index, r := range call.Items {
					response.ImageData.Images[index] = r.Link
				}

				now := time.Now()
				mydb.Insert(keyword, response.ImageData.Images, 1, now, now)
			}
			response.Status = "success"
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}
