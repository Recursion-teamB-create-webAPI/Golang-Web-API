package dao

import (
	"database/sql"
	"encoding/json"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	testErros "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/test"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase_Connect(t *testing.T) {
	type fields struct {
		UseDb *sql.DB
	}
	type args struct {
		env structs.Env
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Successful connection",
			fields: fields{UseDb: nil},
			args: args{env: structs.Env{
				DatabaseName: utils.GetEnvData(constants.BeforeLevel3).DatabaseName,
				MysqlUri:     utils.GetEnvData(constants.BeforeLevel3).MysqlUri,
			}},
		},
		{
			name:   "Failed connection1",
			fields: fields{UseDb: nil},
			args: args{env: structs.Env{
				DatabaseName: utils.GetEnvData(constants.BeforeLevel3).DatabaseName,
				MysqlUri:     "user:password@tcp(localhost:3306)/dbname",
			}},
		},
		{
			name:   "Failed connection2",
			fields: fields{UseDb: nil},
			args: args{env: structs.Env{
				DatabaseName: "MongoDB",
				MysqlUri:     utils.GetEnvData(constants.BeforeLevel3).MysqlUri,
			}},
		},
		{
			name:   "Failed connection3",
			fields: fields{UseDb: nil},
			args: args{env: structs.Env{
				DatabaseName: "MongoDB",
				MysqlUri:     "user:password@tcp(localhost:3306)/dbname",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// sqlmockを使ってモックデータベースと期待する振る舞いを設定
			db, mock, err := sqlmock.New()
			if err != nil {
				log.Println("failed to init db mock")
			}
			defer db.Close()

			// Database.Connectメソッドをテスト
			mydb := &Database{UseDb: db}
			mydb.Connect(tt.args.env)

			// モックの期待した振る舞いがすべて満たされたかを検証
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDatabase_CreateTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		ie := testErros.NewInitMockDbError()
		log.Println(ie.Error())
		return
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`CREATE TABLE IF NOT EXISTS Images (
        id INT AUTO_INCREMENT PRIMARY KEY,
        item VARCHAR(255) NOT NULL,
        images JSON,
        search_count INT DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`)).WillReturnResult(sqlmock.NewResult(0, 0))

	database := &Database{UseDb: db}
	database.CreateTable()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDatabase_InsertInitData(t *testing.T) {
	t.Run(
		"No insert because data already exists",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				log.Println("failed to init db mock")
			}
			defer db.Close()

			initImi := utils.GetInitImagesJson(constants.BeforeLevel3)

			for i := 0; i < constants.ItemCount; i++ {
				count := 0
				date := "2024-03-05 12:03:06"

				item := initImi.ImageItems[i].Item
				images := initImi.ImageItems[i].ImageData.Images

				// images配列をJSON文字列にエンコード
				imagesJSON, err := json.Marshal(images)
				if err != nil {
					t.Fatalf("failed to marshal images: %v", err)
				}

				// Execの呼び出しを期待する設定
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Images WHERE item = ?;`)).
					WithArgs(item).
					WillReturnRows(sqlmock.NewRows([]string{"id", "item", "images", "search_count", "created_at", "updated_at"}).
						AddRow(i+1, item, string(imagesJSON), count, date, date))
			}

			database := &Database{UseDb: db}
			database.InsertInitData(constants.BeforeLevel3)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		},
	)

	t.Run(
		"Successful insert init data",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				log.Println("failed to init db mock")
			}
			defer db.Close()

			initImi := utils.GetInitImagesJson(constants.BeforeLevel3)

			for i := 0; i < constants.ItemCount; i++ {
				item := initImi.ImageItems[i].Item
				images := initImi.ImageItems[i].ImageData.Images

				// images配列をJSON文字列にエンコード
				imagesJSON, err := json.Marshal(images)
				if err != nil {
					t.Fatalf("failed to marshal images: %v", err)
				}

				// Execの呼び出しを期待する設定
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Images WHERE item = ?;`)).
					WithArgs(item).
					WillReturnRows(sqlmock.NewRows([]string{"id", "item", "images", "search_count", "created_at", "updated_at"}))

				// Prepareの呼び出しを期待する設定
				prep := mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO Images (item, images, search_count) VALUES (?, ?, ?);`))

				// Execの呼び出しを期待する設定
				prep.ExpectExec().
					WithArgs(item, imagesJSON, 0).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}

			database := &Database{UseDb: db}
			database.InsertInitData(constants.BeforeLevel3)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		},
	)

}

func TestDatabase_Find(t *testing.T) {
	t.Run(
		"Successful Find data",
		func(t *testing.T) {
			id := 1
			count := 0
			date := "2024-03-05 12:03:06"

			db, mock, err := sqlmock.New()
			if err != nil {
				log.Println("failed to init db mock")
			}
			defer db.Close()

			initImi := utils.GetInitImagesJson(constants.BeforeLevel3)

			item := initImi.ImageItems[0].Item
			images := initImi.ImageItems[0].ImageData.Images

			// images配列をJSON文字列にエンコード
			imagesJSON, err := json.Marshal(images)
			if err != nil {
				t.Fatalf("failed to marshal images: %v", err)
			}

			// Execの呼び出しを期待する設定
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Images WHERE item = ?;`)).
				WithArgs(item).
				WillReturnRows(sqlmock.NewRows([]string{"id", "item", "images", "search_count", "created_at", "updated_at"}).AddRow(id, item, string(imagesJSON), count, date, date))

			database := &Database{UseDb: db}
			success, _ := database.Find(item)

			if !success {
				t.Errorf("An error occurred!")
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		},
	)

	t.Run(
		"Failed Read data",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				log.Println("failed to init db mock")
			}
			defer db.Close()

			initImi := utils.GetInitImagesJson(constants.BeforeLevel3)

			item := initImi.ImageItems[0].Item

			// Execの呼び出しを期待する設定
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Images WHERE item = ?;`)).
				WithArgs(item).
				WillReturnRows(sqlmock.NewRows([]string{"id", "item", "images", "search_count", "created_at", "updated_at"}))

			database := &Database{UseDb: db}
			success, _ := database.Find(item)

			if success {
				t.Errorf("An error occurred!")
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		},
	)
}

func TestDatabase_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println("failed to init db mock")
	}
	defer db.Close()

	initImi := utils.GetInitImagesJson(constants.BeforeLevel3)

	item := initImi.ImageItems[1].Item
	images := initImi.ImageItems[1].ImageData.Images
	searchCount := 2

	// images配列をJSON文字列にエンコード
	imagesJSON, err := json.Marshal(images)
	if err != nil {
		t.Fatalf("failed to marshal images: %v", err)
	}

	// Prepareの呼び出しを期待する設定
	prep := mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO Images (item, images, search_count) VALUES (?, ?, ?);`))

	// Execの呼び出しを期待する設定
	prep.ExpectExec().
		WithArgs(item, imagesJSON, searchCount).
		WillReturnResult(sqlmock.NewResult(1, 1))

	database := &Database{UseDb: db}
	database.Insert(item, images, searchCount)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDatabase_Update(t *testing.T) {
	item := "cat"

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println("failed to init db mock")
	}
	defer db.Close()

	// Prepareの呼び出しを期待する設定
	prep := mock.ExpectPrepare(regexp.QuoteMeta(`UPDATE Images SET search_count = search_count + 1 WHERE item = ?;`))

	// Execの呼び出しを期待する設定
	prep.ExpectExec().
		WithArgs(item).
		WillReturnResult(sqlmock.NewResult(1, 1))

	database := &Database{UseDb: db}
	database.Update(item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
