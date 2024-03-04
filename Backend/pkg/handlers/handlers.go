package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func SearchHandler(env structs.Env, mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseSearch
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			log.Println("Data could not be retrieved because query parameters were not set")
			response.Status = "failed"
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(env, img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				mydb.Update(env, keyword)
				log.Println(img.Images[0])
				response.ImageData.Images = img.Images
				response.Status = "success"
			} else {
				jsonData, err := os.ReadFile(env.KeyFileName)
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
					response.Status = "failed"
				} else {
					for index, r := range call.Items {
						response.ImageData.Images[index] = r.Link
					}
					mydb.Insert(env, keyword, response.ImageData.Images, constants.SearchInitCount)
					response.Status = "success"
				}
			}
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}

func DescriptionHandler(env structs.Env, mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseDescription
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			log.Println("Data could not be retrieved because query parameters were not set")
			response.Status = "failed"
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(env, img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				response.Description = img
				response.Status = "success"
			} else {
				response.Status = "failed"
			}
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}
