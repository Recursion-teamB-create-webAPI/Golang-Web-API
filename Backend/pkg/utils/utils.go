package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/joho/godotenv"
)

func GetEnvData() structs.Env {
	envPath := GetWalkTargetPath(".env")

	// envファイルのパスを渡す
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Error loading .env file")
	}

	// .envから値を取得する
	return structs.Env{
		SearchEngineId: os.Getenv("SEARCH_ENGINE_ID"),
		KeyFileName:    os.Getenv("KEY_FILE_NAME"),
		CsePath:        os.Getenv("CSE_PATH"),
		PortNumber:     os.Getenv("PORT_NUMBER"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		TableName:      os.Getenv("TABELE_NAME"),
		MysqlUri:       os.Getenv("MYSQL_URI"),
	}
}

func GetWalkTargetPath(targetFile string) string {
	// カレントディレクトリを示す絶対パスを取得する
	currDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
	}

	// カレントディレクトリの1階層前のパスを取得する
	rootPath := filepath.Dir(currDir)
	targetPath := ""

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
