package dao

import (
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	testErros "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/test"
)

func TestCreateUsersTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		ie := testErros.NewInitMockDbError()
		log.Println(ie.Error())
		return
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`CREATE TABLE IF NOT EXISTS Users (
		id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL,
        password VARCHAR(100) NOT NULL
        );`)).WillReturnResult(sqlmock.NewResult(0, 0))

	database := &Database{UseDb : db}
	database.CreateTable()
	err = CreateUsersTable(database.UseDb)

	if err != nil {
		t.Errorf("Failed to create Users table.")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}