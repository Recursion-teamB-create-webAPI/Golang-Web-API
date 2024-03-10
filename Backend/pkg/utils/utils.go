package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func GetEnvData(beforeLevel int) structs.Env {
	envPath := GetWalkTargetPath(".env", beforeLevel)

	// envファイルのパスを渡す
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Error loading .env file")
		return structs.Env{}
	}

	// .envから値を取得する
	return structs.Env{
		SearchEngineId: os.Getenv("SEARCH_ENGINE_ID"),
		KeyFileName:    os.Getenv("KEY_FILE_NAME"),
		CsePath:        os.Getenv("CSE_PATH"),
		PortNumber:     os.Getenv("PORT_NUMBER"),
		MysqlUri:       os.Getenv("MYSQL_URI"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		JwtSecretKey:   os.Getenv("JWT_SECRET_KEY"),
	}
}

func GetInitImagesJson(beforeLevel int) *structs.InitImageItems {
	// JSONファイルを開き読み込む
	jsonPath := GetWalkTargetPath("initImages.json", beforeLevel)
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}

	var initImi structs.InitImageItems

	err = json.Unmarshal(jsonData, &initImi)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &initImi
}

func GetGoogleCustomSearchApiResponse(env structs.Env, keyword string, beforeLevel int) *customsearch.Search {
	keyFilePath := GetWalkTargetPath(env.KeyFileName, beforeLevel)
	jsonData, err := os.ReadFile(keyFilePath)
	if err != nil {
		log.Println(err)
		return nil
	}

	conf, err := google.JWTConfigFromJSON(jsonData, env.CsePath)
	if err != nil {
		log.Println(err)
		return nil
	}

	client := conf.Client(context.Background())
	cseService, err := customsearch.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
		return nil
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
		log.Println(err)
		return nil
	}
	return call
}

func GetWalkTargetPath(targetFile string, beforeLevel int) string {
	// カレントディレクトリを示す絶対パスを取得する
	currDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	rootPath := currDir
	targetPath := ""

	for i := 0; i < beforeLevel; i++ {
		rootPath = filepath.Dir(rootPath)
	}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Base(path) == targetFile {
			targetPath = path
			return nil
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	if targetPath == "" {
		log.Println("Could not find targetFile")
	}
	return targetPath
}

func GetSqlQury(item string, queryArr structs.TotalResultQueryArray) (string, []interface{}) {
	var sqlQuery string
	var orderBy string
	var args []interface{}
	skipRows := (queryArr.Page - 1) * queryArr.PerPage
	baseQuery := "SELECT item, search_count, updated_at FROM Images"
	orderClause := " ORDER BY %s DESC LIMIT ? OFFSET ?;"

	// クエリパラメータの値にあわせて、SQLのクエリをセットする
	if queryArr.Order == "count" {
		orderBy = "search_count"
	} else {
		orderBy = "updated_at"
	}

	if item == "" {
		sqlQuery = fmt.Sprintf(baseQuery+orderClause, orderBy)
		args = append(args, queryArr.PerPage, skipRows)
	} else {
		whereClause := " WHERE item LIKE ?"
		sqlQuery = fmt.Sprintf(baseQuery+whereClause+orderClause, orderBy)
		args = append(args, "%"+item+"%", queryArr.PerPage, skipRows)
	}
	return sqlQuery, args
}

func QueryParameterCheck(page string, perPage string, order string) (bool, structs.TotalResultQueryArray) {
	var queryArr structs.TotalResultQueryArray

	// デフォルト値を設定
	page, perPage, order = SetDefaultValue(page, perPage, order)

	// クエリパラメータのチェック
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		log.Println(constants.ErrMessageQueryNotCorrect)
		return false, queryArr
	}

	perPageNum, err := strconv.Atoi(perPage)
	if err != nil {
		log.Println(constants.ErrMessageQueryNotCorrect)
		return false, queryArr
	}

	pageSuccess := pageNum >= 1 && pageNum <= 10
	perPageSuccess := perPageNum >= 1 && perPageNum <= 10
	orderSuccess := order == "count" || order == "newest"
	success := pageSuccess && perPageSuccess && orderSuccess

	queryArr = structs.TotalResultQueryArray{
		Page:    pageNum,
		PerPage: perPageNum,
		Order:   order,
	}
	return success, queryArr
}

func SetDefaultValue(page string, perPage string, order string) (string, string, string) {
	// クエリパラメータが設定されていない場合のデフォルト値設定
	if page == "" {
		page = "1"
	}

	if perPage == "" {
		perPage = "5"
	}

	if order == "" {
		order = "count"
	}
	return page, perPage, order
}
