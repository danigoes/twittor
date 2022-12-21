package routers

import (
	"net/http"
	"github.com/danigoes/twittor/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "Error trying to delete tweet " + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}