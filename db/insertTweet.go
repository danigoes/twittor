package db

import (
	"context"
	"time"
	"github.com/danigoes/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx, cancel :=  context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	record := bson.M {
		"userid": t.UserID,
		"message": t.Message,
		"date": t.Date,
	}
	result, err := col.InsertOne(ctx, record)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}