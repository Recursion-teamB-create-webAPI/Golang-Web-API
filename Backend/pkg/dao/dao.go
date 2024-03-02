package dao

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	UseDb  *mongo.Database
	Col    *mongo.Collection
}

func (db *Database) Connect(env structs.Env, ctx context.Context) {
	opt := options.Client().ApplyURI(env.MongoUri)
	err := opt.Validate()
	if err != nil {
		log.Println(err)
	}
	db.Client, err = mongo.Connect(ctx, opt)
	if err != nil {
		log.Println(err)
	}
}

func (db *Database) Disconnect(ctx context.Context) error {
	return db.Client.Disconnect(ctx)
}

func (db *Database) Ping(ctx context.Context) {
	err := db.Client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Ping to MongoDB server succeeded")
}

func (db *Database) SetDbCol(env structs.Env) {
	db.UseDb = db.Client.Database(env.MongoDatabase)
	db.Col = db.UseDb.Collection(env.MongoCollection)
}

func (db *Database) InsertInitData() {
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

	var imi structs.InitImageItems

	json.Unmarshal(jsonData, &imi)

	now := time.Now()

	// ドキュメントの挿入
	for i := 0; i < constants.ItemCount; i++ {
		var img structs.DatabaseImage

		success, _ := db.Find(img, imi.ImageItems[i].Item)
		if !success {
			db.Insert(imi.ImageItems[i].Item, imi.ImageItems[i].ImageData.Images, 0, now, now)
		}
	}
}

func (db *Database) Find(img structs.DatabaseImage, item string) (bool, structs.DatabaseImage) {
	err := db.Col.FindOne(
		context.TODO(),
		bson.D{{Key: "item", Value: item}}).Decode(&img)
	if err != nil {
		log.Println(err)
	}
	return item == img.Item, img
}

func (db *Database) Insert(item string, images [constants.SearchResultNumber]string, searchCount int, createdAt time.Time, updatedAt time.Time) {
	_, err := db.Col.InsertOne(
		context.TODO(),
		bson.D{
			{Key: "item", Value: item},
			{Key: "images", Value: images},
			{Key: "search_count", Value: searchCount},
			{Key: "created_at", Value: createdAt},
			{Key: "updated_at", Value: updatedAt},
		},
	)
	if err != nil {
		log.Println(err)
	}
}

func (db *Database) Update(itemName string) {
	_, err := db.Col.UpdateOne(
		context.TODO(),
		bson.D{
			{Key: "item", Value: itemName},
		},
		bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: "search_count", Value: 1},
			}},
			{Key: "$currentDate", Value: bson.D{
				{Key: "updated_at", Value: true},
			}},
		},
	)
	if err != nil {
		log.Println(err)
	}
}
