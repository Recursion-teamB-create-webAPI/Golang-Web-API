package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/dao"
	dbError "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/db"
	utilError "github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/errors/util"
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


		var resp structs.ResponseSignIn

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			he := utilError.NewHTTPMethodNotAllowedError(r.Method)
			resp.Status = http.StatusMethodNotAllowed
			resp.Username = ""
			resp.Token = ""
			resp.Error = he.Error()
			json.NewEncoder(w).Encode(resp);
			return
		}

		var user structs.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			me := dbError.NewMapToStructError(user)
			log.Println(me.Error())
			resp.Status = http.StatusBadRequest
			resp.Username = ""
			resp.Token = ""
			resp.Error = me.Error()
			json.NewEncoder(w).Encode(resp);
			return
		}

		query := `SELECT username, password FROM Users`
		rows, err := db.UseDb.Query(query)
		if err != nil {
			qe := dbError.NewDbQueryError(query)
			log.Println(qe.Error())
			resp.Status = http.StatusInternalServerError
			resp.Username = ""
			resp.Token = ""
			resp.Error = qe.Error()
			json.NewEncoder(w).Encode(resp);
			return
		}
		defer rows.Close()

		// Search User
		for rows.Next() {
			var u string
			var p string

			err := rows.Scan(&u, &p)
			if err != nil {
				se := dbError.NewDbRowScanError("Users");
				log.Println(se.Error())
				resp.Status = http.StatusInternalServerError
				resp.Username = ""
				resp.Token = ""
				resp.Error = se.Error()
				json.NewEncoder(w).Encode(resp);
				return
			}

			if u == user.Username {
				err = bcrypt.CompareHashAndPassword([]byte(p), []byte(user.Password))
				if err != nil {
					pe := dbError.NewPasswordDoesNotMatchError();
					log.Println(pe.Error())
					resp.Status = http.StatusBadRequest
					resp.Username = ""
					resp.Token = ""
					resp.Error = pe.Error() 
					json.NewEncoder(w).Encode(resp);
					return
				} else {
					//Create JWT token
					tokenString, err := jwt.GenerateToken(env, user.Username)
					if err != nil {
						ge := dbError.NewGenerateJwtTokenError()
						log.Println(ge.Error())
						resp.Status = http.StatusBadRequest
						resp.Username = ""
						resp.Token = ""
						resp.Error = ge.Error()
						json.NewEncoder(w).Encode(resp);
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
					resp.Status = http.StatusOK
					resp.Username = user.Username
					resp.Token = tokenString
					resp.Error = ""
					json.NewEncoder(w).Encode(structs.ResponseSignIn(resp))
					return
				}
			}
		}
		log.Println(dbError.NewFindUserError(resp.Username).Error())
		resp.Status = http.StatusBadRequest
		resp.Username = ""
		resp.Token = ""
		resp.Error = dbError.NewFindUserError(resp.Username).Error()
		json.NewEncoder(w).Encode(resp);
	}
}
