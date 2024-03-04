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
			errMessage := "Data could not be retrieved because query parameters were not set"

			log.Println(errMessage)
			response.Status = "failed"
			response.Cause = errMessage
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(env, img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				mydb.Update(env, keyword)

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
					errMessage := "The image could not be retrieved because the daily usage limit for the Google Custom Search API has been reached."

					log.Println(err)
					response.Status = "failed"
					response.Cause = errMessage
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
			errMessage := "Data could not be retrieved because query parameters were not set"

			log.Println(errMessage)
			response.Status = "failed"
			response.Cause = errMessage
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(env, img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				response.Description = img
				response.Status = "success"
			} else {
				errMessage := "Keyword data could not be displayed because it does not exist in the database"

				response.Status = "failed"
				response.Cause = errMessage
			}
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}
