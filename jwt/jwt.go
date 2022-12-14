package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/danigoes/twittor/models"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("IAmAKey")
	payload := jwt.MapClaims{
		"email": t.Email,
		"name": t.Name,
		"lastName": t.LastName,
		"birdthDate": t.BirthDate,
		"biography": t.Biography,
		"location": t.Location,
		"wedSite": t.WebSite,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}