package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
)


func ListHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseList
		// クエリパラメータを解析する
		query := r.URL.Query()
		keyword := query.Get("keyword")

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