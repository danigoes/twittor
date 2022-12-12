package middlew

import (
	"net/http"
	"github.com/danigoes/twittor/db"
)

/**
	CheckDB middlew to know the DB status
**/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection lost with DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}