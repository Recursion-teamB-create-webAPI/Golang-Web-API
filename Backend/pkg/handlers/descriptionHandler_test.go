package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
)



func TestDescriptionHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println("failed to init db mock")
	}
	defer db.Close()

	// クエリパラメータにキーワードが含まれる場合のテスト
	t.Run("with keyword", func(t *testing.T) {

		initImi := utils.GetInitImagesJson(constants.BeforeLevel3)
		item := initImi.ImageItems[0].Item
		images := initImi.ImageItems[0].ImageData.Images

		id := 1
		count := 0
		date := "2024-03-05 12:03:06"

		// images配列をJSON文字列にエンコード
		imagesJSON, err := json.Marshal(images)
		if err != nil {
			t.Fatalf("failed to marshal images: %v", err)
		}

		// 期待値設定
		wantBody := structs.ResponseDescription{
			Description: structs.DatabaseImage{
				Id:   id,
				Item: item,
				ImageData: structs.ImageArray{
					Images: images,
				},
				SearchCount: count,
				CreatedAt:   date,
				UpdatedAt:   date,
			},
			Status: "success",
			Cause:  "",
		}

		wantBodyJSON, err := json.Marshal(wantBody)
		if err != nil {
			t.Fatalf("failed to marshal images: %v", err)
		}

		wantBodyStr := string(wantBodyJSON)

		// Execの呼び出しを期待する設定(Find関数)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Images WHERE item = ?;`)).
			WithArgs(item).
			WillReturnRows(sqlmock.NewRows([]string{"id", "item", "images", "search_count", "created_at", "updated_at"}).AddRow(id, item, string(imagesJSON), count, date, date))

		database := &dao.Database{UseDb: db}

		req, err := http.NewRequest("GET", "/search?keyword=cat", nil)
		if err != nil {
			t.Fatal(err)
		}

		// レスポンスを受け止める*httptest.ResponseRecorder
		got := httptest.NewRecorder()
		handler := http.HandlerFunc(DescriptionHandler(database))

		handler.ServeHTTP(got, req)

		// Bodyは*bytes.Buffer型なので文字列を比較する
		if got := strings.TrimSpace(got.Body.String()); got != wantBodyStr {
			t.Errorf("want %s, but got %s", wantBodyStr, got)
		}
	})
	// クエリパラメータにキーワードが含まれない場合のテスト
	t.Run("without keyword", func(t *testing.T) {
		// 期待値設定
		wantBody := structs.ResponseDescription{
			Description: structs.DatabaseImage{
				Id:   0,
				Item: "",
				ImageData: structs.ImageArray{
					Images: [constants.SearchResultNumber]string{},
				},
				SearchCount: 0,
				CreatedAt:   "",
				UpdatedAt:   "",
			},
			Status: "failed",
			Cause:  constants.ErrMessageQuery,
		}

		wantBodyJSON, err := json.Marshal(wantBody)
		if err != nil {
			t.Fatalf("failed to marshal images: %v", err)
		}

		wantBodyStr := string(wantBodyJSON)

		database := &dao.Database{UseDb: db}

		req, err := http.NewRequest("GET", "/search?keyword=", nil)
		if err != nil {
			t.Fatal(err)
		}

		// レスポンスを受け止める*httptest.ResponseRecorder
		got := httptest.NewRecorder()
		handler := http.HandlerFunc(DescriptionHandler(database))

		handler.ServeHTTP(got, req)

		// Bodyは*bytes.Buffer型なので文字列を比較する
		if got := strings.TrimSpace(got.Body.String()); got != wantBodyStr {
			t.Errorf("want %s, but got %s", wantBodyStr, got)
		}
	})
}
