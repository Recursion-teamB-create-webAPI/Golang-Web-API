package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/joho/godotenv"
)

func GetEnvData() structs.Env {
	// カレントディレクトリを示す絶対パスを取得する
	currDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
	}

	// constants.RootDirLevel前の階層のパスを取得する
	rootPath := GetbeforeDirPath(currDir, constants.RootDirLevel)
	// .envのパスを再帰で探す
	targetPath := GetWalkTargetPath(rootPath, ".env")

	// envファイルのパスを渡す
	err = godotenv.Load(targetPath)
	if err != nil {
		log.Println("Error loading .env file")
	}

	// .envから値を取得する
	return structs.Env{
		SearchEngineId: os.Getenv("SEARCH_ENGINE_ID"),
		CsePath:        os.Getenv("CSE_PATH"),
		PortNumber:     os.Getenv("PORT_NUMBER"),
	}
}

func GetbeforeDirPath(path string, beforeLevel int) string {

	for i := 0; i < beforeLevel; i++ {
		path = filepath.Dir(path)
	}
	return path
}

func GetWalkTargetPath(startPath string, targetFile string) string {
	targetPath := ""

	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
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
