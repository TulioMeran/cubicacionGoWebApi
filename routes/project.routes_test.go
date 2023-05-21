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

func TestProjectHandlers(t *testing.T) {

	db.DbConnection()

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	wr := httptest.NewRecorder()

	GetProjectsHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 but got %v", res.StatusCode)
	}

	req = httptest.NewRequest(http.MethodPost, "/project", nil)
	wr = httptest.NewRecorder()

	PostProjectHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expect status code 400 but got %v", res.StatusCode)
	}

	bodyRequest := models.Project{
		Title: "BARQUITA",
	}

	body, err := json.Marshal(bodyRequest)

	if err != nil {
		t.Error(err)
	}

	req = httptest.NewRequest(http.MethodPost, "/project", bytes.NewBuffer(body))
	wr = httptest.NewRecorder()

	PostProjectHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expect status code 201 but got %v", res.StatusCode)
	}

	var newProject models.Project

	json.NewDecoder(res.Body).Decode(&newProject)

	if newProject.ID == 0 {
		t.Errorf("Expect Project ID but got %v", newProject.ID)
	}

	bodyRequestPut := models.Project{
		Title: "Prueba 2",
	}

	body, err = json.Marshal(bodyRequestPut)

	if err != nil {
		t.Error(err)
	}

	req = httptest.NewRequest(http.MethodPut, "/project?id="+strconv.Itoa(int(newProject.ID)), bytes.NewBuffer(body))
	wr = httptest.NewRecorder()

	PutProjectHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 but got %v", res.StatusCode)
	}

	req = httptest.NewRequest(http.MethodDelete, "/project?id="+strconv.Itoa(int(newProject.ID)), nil)
	wr = httptest.NewRecorder()

	DeleteProjectHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Expect status code 204 but got %v", res.StatusCode)
	}

}
