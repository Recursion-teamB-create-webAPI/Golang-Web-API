package dao

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	UseDb *sql.DB
}

func (db *Database) Connect(env structs.Env) {
	database, err := sql.Open(env.DatabaseName, env.MysqlUri)
	if err != nil {
		log.Println(err)
		return
	}

	db.UseDb = database
}

func (db *Database) CreateTable() {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS Images (
        id INT AUTO_INCREMENT PRIMARY KEY,
        item VARCHAR(255) NOT NULL,
        images JSON,
        search_count INT DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`

	_, err := db.UseDb.Exec(createTableSQL)
	if err != nil {
		log.Println(err)
	}
}

func (db *Database) InsertInitData(beforeLevel int) {
	initImi := utils.GetInitImagesJson(beforeLevel)

	if initImi != nil {
		for i := 0; i < constants.ItemCount; i++ {
			success, _ := db.Find(initImi.ImageItems[i].Item)
			if !success {
				db.Insert(initImi.ImageItems[i].Item, initImi.ImageItems[i].ImageData.Images, 0)
			}
		}
	}
}

func (db *Database) Find(item string) (bool, structs.DatabaseImage) {
	var img structs.DatabaseImage
	query := `SELECT * FROM Images WHERE item = ?;`
	res, err := db.UseDb.Query(query, string(item))
	if err != nil {
		log.Println(err)
	}

	var imagesJSON string

	for res.Next() {
		err = res.Scan(&img.Id, &img.Item, &imagesJSON, &img.SearchCount, &img.CreatedAt, &img.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
	}

	err = json.Unmarshal([]byte(imagesJSON), &img.ImageData.Images)
	if err != nil {
		log.Println(err)
	}
	return item == img.Item, img
}

func (db *Database) ReadAllItem() []string {
	var items []string

	query := `SELECT item FROM Images;`
	res, err := db.UseDb.Query(query)
	if err != nil {
		log.Println(err)
	}

	for res.Next() {
		var item string
		err = res.Scan(&item)
		if err != nil {
			log.Println(err)
		}

		items = append(items, item)
	}
	return items
}

func (db *Database) ReadPartialMatchItem(item string) (bool, []string) {
	var items []string

	query := `SELECT item FROM Images WHERE item LIKE ?;`
	res, err := db.UseDb.Query(query, "%"+item+"%")
	if err != nil {
		log.Println(err)
	}

	for res.Next() {
		var item string
		err = res.Scan(&item)
		if err != nil {
			log.Println(err)
		}

		items = append(items, item)
	}
	return items != nil, items
}

func (db *Database) ReadTotalResult(item string, queryArr structs.TotalResultQueryArray) (bool, []structs.TotalResultItems) {
	var totalResults []structs.TotalResultItems

	// クエリ実行に必要な値を取得する
	query, args := utils.GetSqlQury(item, queryArr)

	res, err := db.UseDb.Query(query, args...)
	if err != nil {
		log.Println(err)
	}

	for res.Next() {
		var item string
		var searchCount int
		var updatedAt string

		err = res.Scan(&item, &searchCount, &updatedAt)
		if err != nil {
			log.Println(err)
		}

		newTotalResult := structs.TotalResultItems{
			Item:        item,
			SearchCount: searchCount,
			UpdatedAt:   updatedAt,
		}

		totalResults = append(totalResults, newTotalResult)
	}
	return totalResults != nil, totalResults
}

func (db *Database) Insert(item string, images [constants.SearchResultNumber]string, searchCount int) {
	// images配列をJSON文字列にエンコード
	imagesJSON, err := json.Marshal(images)
	if err != nil {
		log.Println(err)
	}

	query := `INSERT INTO Images (item, images, search_count) VALUES (?, ?, ?);`
	ins, err := db.UseDb.Prepare(query)
	if err != nil {
		log.Println(err)
	}

	_, err = ins.Exec(item, imagesJSON, searchCount)
	if err != nil {
		log.Println(err)
	}
}

func (db *Database) Update(itemName string) {
	query := `UPDATE Images SET search_count = search_count + 1 WHERE item = ?;`
	update, err := db.UseDb.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	_, err = update.Exec(itemName)
	if err != nil {
		log.Println(err)
	}
}
