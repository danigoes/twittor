package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTweets struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string `bson:"userid" json:"userId,omitempty"`
	Mensaje string `bson:"message" json:"message,omitempty"`
	Fecha time.Time `bson:"date" json:"date,omitempty"`
}