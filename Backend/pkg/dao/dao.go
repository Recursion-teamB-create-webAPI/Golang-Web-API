package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	UseDb *sql.DB
}

func (db *Database) Connect(env structs.Env) {
	database, err := sql.Open(env.DatabaseName, env.MysqlUri)
	if err != nil {
		log.Println(err)
	}
	db.UseDb = database
}

func (db *Database) CreateTable(env structs.Env) {

	createTableSQL := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        id INT AUTO_INCREMENT PRIMARY KEY,
        item VARCHAR(255) NOT NULL,
        images JSON,
        search_count INT DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`, env.TableName)

	_, err := db.UseDb.Exec(createTableSQL)
	if err != nil {
		log.Println(err)
	}
}

func (db *Database) InsertInitData(env structs.Env) {
	// JSONファイルを開き読み込む
	jsonFile, err := os.Open("initImages.json")
	if err != nil {
		log.Println("Cannot open JSON file", err)
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println("Cannot read JSON data", err)
	}

	var initImi structs.InitImageItems

	json.Unmarshal(jsonData, &initImi)

	// データの挿入
	for i := 0; i < constants.ItemCount; i++ {
		var img structs.DatabaseImage

		success, _ := db.Find(env, img, initImi.ImageItems[i].Item)
		if !success {
			db.Insert(env, initImi.ImageItems[i].Item, initImi.ImageItems[i].ImageData.Images, 0)
		}
	}
}

func (db *Database) Find(env structs.Env, img structs.DatabaseImage, item string) (bool, structs.DatabaseImage) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE item = ?;`, env.TableName)
	res, err := db.UseDb.Query(query, string(item))
	if err != nil {
		log.Println(err)
	}

	var imagesJSON string

	for res.Next() {
		res.Scan(&img.Id, &img.Item, &imagesJSON, &img.SearchCount, &img.CreatedAt, &img.UpdatedAt)
	}

	json.Unmarshal([]byte(imagesJSON), &img.Images)

	if err != nil {
		log.Println(err)
	}
	return item == img.Item, img
}

func (db *Database) Insert(env structs.Env, item string, images [constants.SearchResultNumber]string, searchCount int) {
	// images配列をJSON文字列にエンコード
	imagesJSON, err := json.Marshal(images)
	if err != nil {
		log.Println(err)
	}

	query := fmt.Sprintf(`INSERT INTO %s (item, images, search_count) VALUES (?, ?, ?);`, env.TableName)
	ins, err := db.UseDb.Prepare(query)
	if err != nil {
		log.Println(err)
	}

	// imagesJSONをバイトスライスから文字列に変換してExecに渡す
	_, err = ins.Exec(item, imagesJSON, searchCount)
	if err != nil {
		log.Println(err)
	}
}

func (db *Database) Update(env structs.Env, itemName string) {
	query := fmt.Sprintf(`UPDATE %s SET search_count = search_count + 1 WHERE item = ?;`, env.TableName)
	update, err := db.UseDb.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	_, err = update.Exec(itemName)
	if err != nil {
		log.Println(err)
	}
}
