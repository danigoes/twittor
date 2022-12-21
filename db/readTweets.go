package db

import (
	"context"
	"time"
	"log"
	"github.com/danigoes/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page  int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel :=  context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ReturnTweets

	condition := bson.M {
		"userid": ID,
	}

	optionsFilter := options.Find()
	optionsFilter.SetLimit(20)
	optionsFilter.SetSort(bson.D{{Key:"date", Value: -1}})
	optionsFilter.SetSkip((page - 1) * 20)

	records, err := col.Find(ctx, condition, optionsFilter)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for records.Next(context.TODO()) {
		var record models.ReturnTweets
		err := records.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}

	return results, true
}
