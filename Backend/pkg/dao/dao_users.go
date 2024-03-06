package dao

import (
	"database/sql"
	"fmt"
	"log"

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
		log.Println("Failed to create Users Table")
		return err
	}
	fmt.Println("Success to create Users Table")
	return nil
}

func InsertUser(db *sql.DB, username string, rawPassword string) error {
	hashedPassword, err := encryptPassword(rawPassword)
	if err != nil {
		return err
	}
	fmt.Printf("Hashed Password: %v\n", hashedPassword)
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
		log.Printf("Failed to insert new user into Users table.: %v\n", err.Error())
		return err
	}
	return nil
}

func SearchUser(db *sql.DB, username string, rawPassword string) error {
	query := `
	SELECT username, password FROM Users
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to exec search user query: %v\n", err.Error())
		return err
	}
	defer rows.Close()

	alreadyExists, err := JudgeUserAlreadyExists(rows, username)
	if err != nil {
		return err
	}
	if alreadyExists {
		return NewHaveAlreadyUserError(username)
	}

	return nil
}

func JudgeUserAlreadyExists(rows *sql.Rows, username string) (bool, error) {
	for rows.Next() {
		var u string
		var p string

		err := rows.Scan(&u, &p)
		if err != nil {
			log.Printf("Failed to scan db row: %v\n", err.Error())
			return false, err
		}
		//logging
		fmt.Printf("Username: %v\n", u)
		fmt.Printf("Password: %v\n", p)

		if u == username {
			return true, nil
		}
	}
	return false, nil
}

func encryptPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v\n", err.Error())
		return "Failed to hash password", err
	}
	return string(hashedPassword), nil
}

// After merege, go to errors.go
type HaveAlreadyUserError struct {
	username string
}

func (e *HaveAlreadyUserError) Error() string {
	return fmt.Sprintf("You have already registered. Username: %v\n", e.username)
}

func NewHaveAlreadyUserError(username string) *HaveAlreadyUserError {
	return &HaveAlreadyUserError{
		username: username,
	}
}
