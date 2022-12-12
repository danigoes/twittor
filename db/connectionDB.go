package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://danigoes:gotwittor@cluster0.sipgyir.mongodb.net/?retryWrites=true&w=majority")

/**
	ConnectDB method to connect to DB
**/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) // Call check DB availability
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection successful")
	return client
}

/**
	CheckConnection ping to DB
**/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}