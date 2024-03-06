package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
)

func SearchHandler(env structs.Env, mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseSearch
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			log.Println(constants.ErrMessageQuery)
			response.Status = "failed"
			response.Cause = constants.ErrMessageQuery
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				mydb.Update(keyword)

				response.ImageData.Images = img.ImageData.Images
				response.Status = "success"
			} else {
				call := utils.GetGoogleCustomSearchApiResponse(env, keyword, constants.BeforeLevel0)
				if call != nil {
					for index, r := range call.Items {
						response.ImageData.Images[index] = r.Link
					}
					mydb.Insert(keyword, response.ImageData.Images, constants.SearchInitCount)
					response.Status = "success"
				} else {
					response.Status = "failed"
					response.Cause = constants.ErrMessageApi
				}
			}
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}

func DescriptionHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseDescription
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			log.Println(constants.ErrMessageQuery)
			response.Status = "failed"
			response.Cause = constants.ErrMessageQuery
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(img, keyword)
			// keywordがデータベースに存在するかチェックする
			if success {
				response.Description = img
				response.Status = "success"
			} else {
				response.Status = "failed"
				response.Cause = constants.ErrMessageDb
			}
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}
