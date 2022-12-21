package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/danigoes/twittor/db"
	"github.com/danigoes/twittor/models"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	record := models.SaveTweet {
		UserID: IDUser, 
		Message: message.Message,
		Date: time.Now(),
	}

	_, status, err := db.InsertTweet(record)
	if err != nil {
		http.Error(w, "Error trying to insert tweet " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "We couldn't insert tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}