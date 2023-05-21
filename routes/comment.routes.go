package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	var comments []models.Comment

	result := db.DB.Find(&comments)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting comments: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&comments)

}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Comment
	json.NewDecoder(r.Body).Decode(&t)

	if t.CubicacionID < 1 {
		http.Error(w, "cubicacionid is required", http.StatusBadRequest)
		return
	}

	if t.StatusCubicacionID < 1 {
		http.Error(w, "statuscubicacionid is required", http.StatusBadRequest)
		return
	}

	if len(t.Description) < 1 {
		http.Error(w, "description is required", http.StatusBadRequest)
		return
	}

	t.UserID = UserID

	result := db.DB.Create(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs creating comment: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&t)

}

func PutCommentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.Comment

	json.NewDecoder(r.Body).Decode(&t)

	if t.CubicacionID < 1 {
		http.Error(w, "cubicacionid is required", http.StatusBadRequest)
		return
	}

	if t.StatusCubicacionID < 1 {
		http.Error(w, "statuscubicacionid is required", http.StatusBadRequest)
		return
	}

	if len(t.Description) < 1 {
		http.Error(w, "description is required", http.StatusBadRequest)
		return
	}

	var current models.Comment

	result := db.DB.First(&current, id)

	if current.ID == 0 {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting comment for putComment", http.StatusInternalServerError)
		return
	}

	if current.UserID != UserID {
		http.Error(w, "Comment only can be update for the user who created.", http.StatusBadGateway)
		return
	}

	if current.StatusCubicacionID != t.StatusCubicacionID {
		http.Error(w, "Comment only can be update in the same status.", http.StatusBadGateway)
		return
	}

	current.Description = t.Description

	result = db.DB.Save(&current)

	if result.Error != nil {
		http.Error(w, "Error occurrs saving the comment: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.Comment

	result := db.DB.First(&t, id)

	if t.ID == 0 {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting comment for delete: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.DB.Unscoped().Delete(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs deleting comment: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
