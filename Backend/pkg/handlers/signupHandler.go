package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
)

func SignUpHandler(env structs.Env, db *dao.Database) http.HandlerFunc {
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
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		fmt.Printf("Username: %v\n", user.Username)
		fmt.Printf("Password: %v\n", user.Password)

		err = dao.CreateUsersTable(db.UseDb)
		if err != nil {
			log.Println("Failed to create Users table.")
			http.Error(w, "Failed to create User table.", http.StatusInternalServerError)
			return
		}

		err = dao.InsertUser(db.UseDb, user.Username, user.Password)
		if err != nil {
			log.Println("Failed to insert user into Users table.")
			http.Error(w, "Failed to insert user into Users table.", http.StatusInternalServerError)
			return
		}

		var resp structs.ResponseSignUp
		resp.Username = user.Username

		json.NewEncoder(w).Encode(resp)
	}
}
