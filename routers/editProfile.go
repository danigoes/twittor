package routers

import (
	"encoding/json"
	"net/http"
	"github.com/danigoes/twittor/db"
	"github.com/danigoes/twittor/models"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Wrong data" + err.Error(), 400)
		return
	}

	var status bool
	status, err = db.EditRecord(t, IDUser)
	if err != nil {
		http.Error(w, "Something went wrong trying to edit record. " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Record has not been edit it", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}