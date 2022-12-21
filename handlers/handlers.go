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
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.CheckDB(middlew.ValidateJWT(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}