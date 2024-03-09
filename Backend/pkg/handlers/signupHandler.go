package handlers

import (
	"encoding/json"
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

		var resp structs.ResponseSignUp

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			resp.Status = http.StatusMethodNotAllowed
			resp.Id = -1 
			resp.Username = ""
			json.NewEncoder(w).Encode(resp)
			return
		}

		var user structs.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			resp.Status = http.StatusInternalServerError
			resp.Id = -1
			resp.Username = ""
			json.NewEncoder(w).Encode(resp)
			return
		}

		err = dao.CreateUsersTable(db.UseDb)
		if err != nil {
			log.Println("Failed to create Users table.")
			resp.Status = http.StatusInternalServerError
			resp.Id = -1
			resp.Username = ""
			json.NewEncoder(w).Encode(resp)
			return
		}

		err = dao.InsertUser(db.UseDb, user.Username, user.Password)
		if err != nil {
			log.Println("Failed to insert user into Users table.")
			resp.Status = http.StatusBadRequest
			resp.Id = -1
			resp.Username = ""
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp.Status = http.StatusOK
		resp.Username = user.Username
		json.NewEncoder(w).Encode(resp)
	}
}
