package handlers

import (
	"net/http"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
)

func SignOutHandler(env structs.Env, db *dao.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    "",
			Expires:  time.Now().AddDate(-1, 0, 0), 
			HttpOnly: true,
			Secure:   true, 
			SameSite: http.SameSiteStrictMode,
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out successfully"))
	}
}