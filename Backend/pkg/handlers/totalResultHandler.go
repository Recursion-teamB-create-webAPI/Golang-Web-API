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

func TotalResultHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")


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