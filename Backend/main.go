package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

type Env struct {
	searchEngineId string
	csePath        string
	portNumber     string
}

type ResponseImage struct {
	Images [10]string `json:"images"`
}

func (env *Env) searchHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを解析する
	query := r.URL.Query()
	keyword := query.Get("keyword")

	jsonData, err := os.ReadFile("search-key.json")
	if err != nil {
		log.Println(err)
	}

	conf, err := google.JWTConfigFromJSON(jsonData, env.csePath)
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
	search.Cx(env.searchEngineId)
	// Custom Search Engineで「画像検索」をオンにする
	search.SearchType("image")

	search.Start(1)
	call, err := search.Do()
	if err != nil {
		log.Println(err)
	}

	var response ResponseImage

	for index, r := range call.Items {
		response.Images[index] = r.Link
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func getEnvData() Env {
	// envファイルのパスを渡す
	path := strings.Replace(getAbsolutePath(".env"), "Backend", "Env", -1)
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file")
	}

	// .envから値を取得する
	return Env{
		searchEngineId: os.Getenv("SEARCH_ENGINE_ID"),
		csePath:        os.Getenv("CSE_PATH"),
		portNumber:     os.Getenv("PORT_NUMBER"),
	}
}

func getAbsolutePath(path string) string {
	curentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return filepath.Join(curentDir, path)
}

func main() {
	// .envから値を取得する
	env := getEnvData()

	log.Println("Starting the server!")

	// 検索用エンドポイントにアクセスされたら呼び出す
	http.Handle("/api/search", http.HandlerFunc(env.searchHandler))

	// 8000番ポートでサーバを開始
	http.ListenAndServe(env.portNumber, nil)
}
