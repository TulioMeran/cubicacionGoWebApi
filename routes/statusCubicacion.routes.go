package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/dto"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func GetStatusCubicacionsHandler(w http.ResponseWriter, r *http.Request) {
	var status []models.StatusCubicacion
	result := db.DB.Find(&status)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting statusCubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var statusDto []dto.StatusCubicacion

	for _, s := range status {
		var element = dto.StatusCubicacion{
			Codigo:      int(s.ID),
			Descripcion: s.Description,
		}
		statusDto = append(statusDto, element)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&statusDto)
}

func PostStatusCubicacionHandler(w http.ResponseWriter, r *http.Request) {
	var t models.StatusCubicacion

	json.NewDecoder(r.Body).Decode(&t)

	if len(t.Description) < 1 {
		http.Error(w, "description is required", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs creating StatusCubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&t)

}

func DeleteStatusCubicacion(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("id")) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.StatusCubicacion

	id := r.URL.Query().Get("id")

	result := db.DB.First(&t, id)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting statusCubicacion to delete: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.DB.Unscoped().Delete(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs deleting statusCubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PutStatusCubicacion(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.StatusCubicacion
	json.NewDecoder(r.Body).Decode(&t)

	var status models.StatusCubicacion
	result := db.DB.First(&status, id)

	if status.ID < 1 {
		http.Error(w, "StatusCubicacion not found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting statusCubicacion for putStatusCubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(t.Description) > 1 {
		status.Description = t.Description
	}

	result = db.DB.Save(&status)

	if result.Error != nil {
		http.Error(w, "Error occurrs saving statusCubicacion for putStatusCubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
