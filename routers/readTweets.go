package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/danigoes/twittor/db"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "page parameter is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "page parameter has to be a number grater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, correct := db.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}