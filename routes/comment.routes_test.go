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

func TestGetCommentsHandler(t *testing.T) {

	db.DbConnection()

	req := httptest.NewRequest(http.MethodGet, "/comment", nil)

	wr := httptest.NewRecorder()

	GetCommentsHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v", res.StatusCode)
	}
}

func TestPostCommentHandlerBadRequest(t *testing.T) {

	db.DbConnection()

	req := httptest.NewRequest(http.MethodPost, "/comment", nil)
	wr := httptest.NewRecorder()

	PostCommentHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expect status code 400 but got %v", res.StatusCode)
	}

}

func TestPostPutDeleteCommentHandlerOk(t *testing.T) {
	db.DbConnection()

	bodyRequest := models.Comment{
		Description:        "aloha",
		CubicacionID:       2,
		StatusCubicacionID: 1,
	}

	body, err := json.Marshal(bodyRequest)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/comment", bytes.NewBuffer(body))
	wr := httptest.NewRecorder()

	PostCommentHandler(wr, req)

	res := wr.Result()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expect status code 201 but got %v", res.StatusCode)
	}

	var newComment models.Comment

	json.NewDecoder(res.Body).Decode(&newComment)

	if newComment.ID < 1 {
		t.Errorf("Expect comment ID to be present but got %v", newComment.ID)
	}

	UserID = newComment.UserID

	bodyRequestPut := models.Comment{
		Description:        "klok",
		CubicacionID:       newComment.CubicacionID,
		StatusCubicacionID: newComment.StatusCubicacionID,
		UserID:             newComment.UserID,
	}

	body, err = json.Marshal(bodyRequestPut)

	if err != nil {
		t.Error(err)
	}

	id := int(newComment.ID)

	ruta := "/comment?id=" + strconv.Itoa(id)

	req = httptest.NewRequest(http.MethodPut, ruta, bytes.NewBuffer(body))
	wr = httptest.NewRecorder()

	PutCommentHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 but got %v", res.StatusCode)
	}

	req = httptest.NewRequest(http.MethodDelete, "/comment?id="+strconv.Itoa(int(newComment.ID)), nil)
	wr = httptest.NewRecorder()

	DeleteCommentHandler(wr, req)

	res = wr.Result()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Expect status code 204 but got %v", res.StatusCode)
	}

}
