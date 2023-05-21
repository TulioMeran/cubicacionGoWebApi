package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func TestGetCubicacionHandler(t *testing.T) {

	db.DbConnection()

	req := httptest.NewRequest(http.MethodGet, "/cubicacion", nil)
	wr := httptest.NewRecorder()

	GetCubicacionesHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 but got %v", res.StatusCode)
	}
}

func TestPostPutDeleteCubicacionHandler(t *testing.T) {

	db.DbConnection()

	bodyRequest := models.Cubicacion{
		Description:        "Cubicacion prueba 1",
		ProjectID:          5,
		StatusCubicacionID: 1,
	}

	body, err := json.Marshal(bodyRequest)

	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/cubicacion", bytes.NewBuffer(body))
	wr := httptest.NewRecorder()

	PostCubicacionesHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expect status code 201 but got %v", res.StatusCode)
	}

	var newCubicacion models.Cubicacion

	json.NewDecoder(res.Body).Decode(&newCubicacion)

	id := int(newCubicacion.ID)

	var bodyRequetPut = models.Cubicacion{
		Description:        "Prueba",
		StatusCubicacionID: newCubicacion.StatusCubicacionID,
		ProjectID:          newCubicacion.ProjectID,
		Observation:        "Prueba",
	}

	body, err = json.Marshal(bodyRequetPut)

	if err != nil {
		t.Error(err)
	}

	req = httptest.NewRequest(http.MethodPut, "/cubicacion?id="+strconv.Itoa(id), bytes.NewBuffer(body))
	wr = httptest.NewRecorder()

	PutCubicacionHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 but got %v", res.StatusCode)
	}

	req = httptest.NewRequest(http.MethodDelete, "/cubicacion?id="+strconv.Itoa(id), nil)
	wr = httptest.NewRecorder()

	DeleteCubicacionHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Expect status code 204 but got %v", res.StatusCode)
	}

}
