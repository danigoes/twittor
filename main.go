package main

import (
	"log"
	"github.com/danigoes/twittor/handlers"
	"github.com/danigoes/twittor/db"
)

func main()  {
	if db.CheckConnection() == 0 {
		log.Fatal("No Connection")
		return
	}
	handlers.Handlers()
}