package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/jwt"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(env structs.Env, db *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user structs.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		query := `SELECT username, password FROM Users`
		rows, err := db.UseDb.Query(query)
		if err != nil {
			log.Printf("Failed to exec db query: %v\n", query)
			http.Error(w, "Failed to exec db select query", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Search User
		for rows.Next() {
			var u string
			var p string

			err := rows.Scan(&u, &p)
			if err != nil {
				log.Printf("Failed to scan db row: %v\n", err.Error())
				http.Error(w, "Failed to scan db row", http.StatusInternalServerError)
			}
			fmt.Printf("Username: %v\n", u)
			fmt.Printf("Password: %v\n", p)

			if u == user.Username {
				err = bcrypt.CompareHashAndPassword([]byte(p), []byte(user.Password))
				if err != nil {
					log.Println("Password does not match")
					http.Error(w, "Password does not match", http.StatusBadRequest)
					return
				}
			}
		}

		//Create JWT token
		tokenString, err := jwt.GenerateToken(env, user.Username)
		if err != nil {
			http.Error(w, "Failed to generate token string.", http.StatusInternalServerError)
			return
		}
		//Set token into cookie
		cookie := http.Cookie{
			Name:     "jwt_token",
			Value:    tokenString,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
			Path:     "/", /*Path should be modified into /username*/
		}
		//return response
		http.SetCookie(w, &cookie)
		var resp structs.ResponseSignIn
		resp.Username = user.Username
		json.NewEncoder(w).Encode(structs.ResponseSignIn(resp))
	}
}
