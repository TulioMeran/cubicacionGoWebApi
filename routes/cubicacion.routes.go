package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func GetCubicacionesHandler(w http.ResponseWriter, r *http.Request) {
	var cubicaciones []models.Cubicacion

	result := db.DB.Preload("Project").Preload("StatusCubicacion").Find(&cubicaciones)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting cubicaciones: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&cubicaciones)

}

func PostCubicacionesHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Cubicacion

	json.NewDecoder(r.Body).Decode(&t)

	if len(t.Description) < 1 {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	if t.ProjectID < 1 {
		http.Error(w, "ProjectID is required", http.StatusBadRequest)
		return
	}

	if t.StatusCubicacionID < 1 {
		http.Error(w, "StatusCubicacionID is required", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs creating cubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&t)

}

func DeleteCubicacionHandler(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("id")) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.Cubicacion
	id := r.URL.Query().Get("id")

	result := db.DB.First(&t, id)

	if result.Error != nil {
		http.Error(w, "Error occurrs finding the cubicacion for delete: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs deleting the cubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
