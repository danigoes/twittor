package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/danigoes/twittor/middlew"
	"github.com/danigoes/twittor/routers"
	"github.com/rs/cors"
)

/**
	Handlers method to set port and listen server
**/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}