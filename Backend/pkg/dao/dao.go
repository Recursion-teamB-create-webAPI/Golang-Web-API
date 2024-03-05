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
			var img structs.DatabaseImage

			success, _ := db.Find(img, initImi.ImageItems[i].Item)
			if !success {
				db.Insert(initImi.ImageItems[i].Item, initImi.ImageItems[i].ImageData.Images, 0)
			}
		}
	}
}

func (db *Database) Find(img structs.DatabaseImage, item string) (bool, structs.DatabaseImage) {
	query := `SELECT * FROM Images WHERE item = ?;`
	res, err := db.UseDb.Query(query, string(item))
	if err != nil {
		log.Println(err)
	}

	var imagesJSON string

	for res.Next() {
		res.Scan(&img.Id, &img.Item, &imagesJSON, &img.SearchCount, &img.CreatedAt, &img.UpdatedAt)
	}

	json.Unmarshal([]byte(imagesJSON), &img.ImageData.Images)

	if err != nil {
		log.Println(err)
	}
	return item == img.Item, img
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
