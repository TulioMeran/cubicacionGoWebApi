package jwt

import (
	"time"

	"github.com/TulioMeran/cubicacionGoWebApi/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func Generator(t models.User) (string, error) {
	myKey := []byte("this_is_a_fucking_key")

	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"lastname": t.LastName,
		"id":       t.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
