package routers

import (
	"encoding/json"
	"net/http"
	"github.com/danigoes/twittor/db"
	"github.com/danigoes/twittor/models"
)

/**
	Register method to create user in DB
**/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Received data error " + err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password must have more than 6 characters", 400)
		return
	}

	_, userFound, _ := db.CheckUserExist(t.Email)
	if userFound == true {
		http.Error(w, "User already exists", 400)
		return
	}

	_, status, err := db.InsertRecord(t)
	if err != nil {
		http.Error(w, "Error trying to insert record " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Failed to insert record", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}