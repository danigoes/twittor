package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/danigoes/twittor/db"
	"github.com/danigoes/twittor/jwt"
	"github.com/danigoes/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or password invalid " + err.Error(), 400)
		return
	}
	if len(t.Email)==0 {
		http.Error(w, "User required", 400)
		return
	}
	document, exist := db.TryLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "User or password invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error trying to generate token " + err.Error(), 400)
		return
	}

	resp := models.LoginResponse {
		Token : jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie {
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}