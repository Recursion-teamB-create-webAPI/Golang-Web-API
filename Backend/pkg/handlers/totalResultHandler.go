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

func TotalResultHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseTotalResult
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")
		page := query.Get("page")
		perPage := query.Get("perpage")
		order := query.Get("order")

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