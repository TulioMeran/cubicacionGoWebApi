package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func TestLoginHandlerBadRequest(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/auth/login", nil)
	req.Header.Set("Content-Type", "application/json")

	wr := httptest.NewRecorder()

	LoginHandler(wr, req)

	res := wr.Result()

	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Expected error to be nil got %v", err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expect status code 400 but got %v", res.StatusCode)
	}
}

func TestLoginHandlerOkRequest(t *testing.T) {

	db.DbConnection()

	bodyRequest := models.Login{
		Email:    "rtulio007@gmail.com",
		Password: "123",
	}

	body, err := json.Marshal(bodyRequest)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer([]byte(body)))

	wr := httptest.NewRecorder()

	LoginHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 but got %v", res.StatusCode)
	}

	defer res.Body.Close()

	var response models.LoginResponse

	json.NewDecoder(res.Body).Decode(&response)

	if len(response.Token) < 1 {
		t.Errorf("Expect to return token but got %v", response.Token)
	}

}
