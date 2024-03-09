package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	searchError "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/search"
	utilError "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/util"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
)

func SearchHandler(env structs.Env, mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		var response structs.ResponseSearch
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			nke := utilError.NewNoKeywordError()
			log.Println(nke.Error())
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
					nge := searchError.NewNoGoogleCustomSearchApiResponseError()
					log.Println(nge.Error())
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


