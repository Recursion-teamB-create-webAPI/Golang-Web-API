package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
		keyword := strings.ToLower(query.Get("keyword"))

		if keyword == "" {
			log.Println(constants.ErrMessageQuery)
			response.Status = "failed"
			response.Cause = constants.ErrMessageQuery
		} else {
			// keywordがデータベースに存在するかチェックする
			success, img := mydb.Find(keyword)

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
		keyword := strings.ToLower(query.Get("keyword"))

		if keyword == "" {
			log.Println(constants.ErrMessageQuery)
			response.Status = "failed"
			response.Cause = constants.ErrMessageQuery
		} else {
			// keywordがデータベースに存在するかチェックする
			success, img := mydb.Find(keyword)

			if success {
				response.Description = img
				response.Status = "success"
			} else {
				log.Println(constants.ErrMessageDb)
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

func ListHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseList
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := strings.ToLower(query.Get("keyword"))

		if keyword == "" {
			response.List = mydb.ReadAllItem()
			response.Status = "success"
		} else {
			// keywordがデータベースに存在するかチェックする
			success, list := mydb.ReadPartialMatchItem(keyword)

			if success {
				response.List = list
				response.Status = "success"
			} else {
				log.Println(constants.ErrMessageNotExist)
				response.Status = "failed"
				response.Cause = constants.ErrMessageNotExist
			}
		}
		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}

func TotalResultHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseTotalResult
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := strings.ToLower(query.Get("keyword"))
		page := strings.ToLower(query.Get("page"))
		perPage := strings.ToLower(query.Get("perpage"))
		order := strings.ToLower(query.Get("order"))

		// クエリパラメータに適切な値がセットされているかのチェック
		successCheck, queryArr := utils.QueryParameterCheck(page, perPage, order)

		if successCheck {
			successRead, totalResults := mydb.ReadTotalResult(keyword, queryArr)

			if successRead {
				response.TotalResult = totalResults
				response.Status = "success"
			} else {
				log.Println(constants.ErrMessageQuerySetValue)
				response.Status = "failed"
				response.Cause = constants.ErrMessageQuerySetValue
			}
		} else {
			log.Println(constants.ErrMessageQueryNotCorrect)
			response.Status = "failed"
			response.Cause = constants.ErrMessageQueryNotCorrect
		}

		// Content-Typeヘッダーをapplication/jsonに設定
		w.Header().Set("Content-Type", "application/json")

		// マップをJSONにエンコードしてレスポンスとして送信
		json.NewEncoder(w).Encode(response)
	}
}
