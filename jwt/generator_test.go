package jwt

import (
	"testing"

	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func TestGenerator(t *testing.T) {
	var user models.User = models.User{
		Email:    "rtulio@gmail.com",
		Name:     "Rafael",
		LastName: "Meran",
		Password: "123456",
	}

	token, err := Generator(user)

	if err != nil {
		t.Error("Token creation is failing")
	}

	if len(token) < 1 {
		t.Errorf("Should create a valid token but it is creating %s instead", token)
	}

}
