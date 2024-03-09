package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	utilError "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/util"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
)



func DescriptionHandler(mydb *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response structs.ResponseDescription
		query := r.URL.Query()
		keyword := query.Get("keyword")

		if keyword == "" {
			nke := utilError.NewNoKeywordError()
			log.Println(nke.Error())
			response.Status = "failed"
			response.Cause = constants.ErrMessageQuery
		} else {
			var img structs.DatabaseImage
			success, img := mydb.Find(keyword)
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
