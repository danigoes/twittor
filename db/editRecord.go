package db

import (
	"context"
	"time"

	"github.com/danigoes/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditRecord(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	// crear el registro que se quiere modificar con solo los campos que tienen cambios
	record := make(map[string]interface{})
	if len(u.Name) > 0 {
		record["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		record["lastName"] = u.LastName
	}
	record["birthDate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		record["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		record["webSite"] = u.WebSite
	}

	// registro de actualizacion
	updtString := bson.M{
		"$set": record,
	}

	// convertir nuestro ID en objID
	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": bson.M{"$eq":objID}}

	_, err := col.UpdateOne(ctx, condition, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}