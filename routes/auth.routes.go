package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/jwt"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
	jwt_go "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var UserID int
var UserEmail string

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Login

	json.NewDecoder(r.Body).Decode(&t)

	if len(t.Email) < 1 {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 1 {
		http.Error(w, "password is required", http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", t.Email).First(&user)

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting user in login", http.StatusInternalServerError)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(t.Password))

	if err != nil {
		http.Error(w, "User or password invalid", http.StatusBadRequest)
		return
	}

	token, err := jwt.Generator(user)

	var resp = models.LoginResponse{
		Name:     user.Name,
		LastName: user.LastName,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("this_is_a_fucking_key")

	claims := &models.Claim{}
	splitToken := strings.Split(tk, " ")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Token invalid format.")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt_go.ParseWithClaims(tk, claims, func(t *jwt_go.Token) (interface{}, error) {
		return myKey, nil
	})

	var user models.User

	if err == nil {
		result := db.DB.Where("email = ?", claims.Email).First(&user)
		if result.Error != nil {
			return claims, false, "", result.Error
		}
		UserID = int(user.ID)
		UserEmail = user.Email
		return claims, true, user.Email, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token invalid")
	}

	return claims, false, "", err

}
