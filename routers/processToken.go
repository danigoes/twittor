package routers

import (
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/danigoes/twittor/db"
	"github.com/danigoes/twittor/models"
)

var Email string
var IDUser string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("IAmAKey")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid") 
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return myKey, nil
	})
	if err == nil {
		_, found, _ := db.CheckUserExist(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalid")
	}
	
	return claims, false, string(""), err
}