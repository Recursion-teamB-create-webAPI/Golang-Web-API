package dao

import (
	"database/sql"
	"fmt"
	"log"

	dbError "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/db"
	"golang.org/x/crypto/bcrypt"
)

func CreateUsersTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS Users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(50) NOT NULL,
            password VARCHAR(100) NOT NULL
        )
    `
	_, err := db.Exec(query)
	if err != nil {
		ce := dbError.NewCreateTableError("Users", query)
		log.Println(ce.Error())
		return ce 
	}
	fmt.Println("Success to create Users Table")
	return nil
}

func InsertUser(db *sql.DB, username string, rawPassword string) error {
	hashedPassword, err := encryptPassword(rawPassword)
	if err != nil {
		return err
	}
	/* If we have that same user (same username and password), we should return error */
	err = SearchUser(db, username, rawPassword)
	if err != nil {
		return err
	}
	query := `
	INSERT INTO Users (username, password)
	VALUES (?, ?)
	`
	password := hashedPassword
	_, err = db.Exec(query, username, password)
	if err != nil {
		ie := dbError.NewInsertUserError(username)
		log.Println(ie.Error())
		return ie 
	}
	return nil
}

func SearchUser(db *sql.DB, username string, rawPassword string) error {
	query := `
	SELECT username, password FROM Users
	`
	rows, err := db.Query(query)
	if err != nil {
		ee := dbError.NewDbQueryError(query)
		log.Println(ee.Error())
		return ee 
	}
	defer rows.Close()

	alreadyExists, err := JudgeUserAlreadyExists(rows, username)
	if err != nil {
		return err
	}
	if alreadyExists {
		return dbError.NewUserAlreadyExistsError(username)
	}

	return nil
}

func JudgeUserAlreadyExists(rows *sql.Rows, username string) (bool, error) {
	for rows.Next() {
		var u string
		var p string

		err := rows.Scan(&u, &p)
		if err != nil {
			de := dbError.NewDbRowScanError("Users")
			log.Println(de.Error())
			return false, de 
		}

		if u == username {
			return true, nil
		}
	}
	return false, nil
}

func encryptPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		ee := dbError.NewEncryptPasswordError();
		log.Printf("Failed to hash password: %v\n", ee.Error())
		return "Failed to hash password", ee
	}
	return string(hashedPassword), nil
}