package routes

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
	"github.com/google/uuid"
)

func GetPathFilesByCubicationId(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	cubicacionId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Id must be a number", http.StatusBadRequest)
		return
	}

	var PathFiles []models.PathFile

	result := db.DB.Where(&models.PathFile{CubicacionID: cubicacionId}).Find(&PathFiles)

	if result.Error != nil {
		http.Error(w, "Error getting pathFiles from database: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&PathFiles)
	return

}

func PostPathFilesByCubicacionId(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("archivo")

	if err != nil {
		http.Error(w, "archivo is required: "+err.Error(), http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.Cubicacion

	result := db.DB.First(&t, id)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting cubicacion in upload file: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if t.ID == 0 {
		http.Error(w, "Cubicacion not found", http.StatusNotFound)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]

	file_name := uuid.New().String()

	var archivo string = "uploads/cubicacion/" + file_name + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error occurrs uploading the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error occurrs copying the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	cubicacionId, _ := strconv.Atoi(id)

	result = db.DB.Model(&models.PathFile{}).Where(&models.PathFile{CubicacionID: cubicacionId}).Update("Active", false)

	if result.Error != nil {
		http.Error(w, "Error occurrs updating old pathFiles: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var newPathFile models.PathFile

	newPathFile.Name = file_name + "." + extension
	newPathFile.Active = true
	newPathFile.CubicacionID = int(t.ID)

	result = db.DB.Create(&newPathFile)

	if result.Error != nil {
		http.Error(w, "Error occurrs creating pathFile: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutPathFilesByCubicationId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	cubicacionId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	var p models.PathFile

	json.NewDecoder(r.Body).Decode(&p)

	var pathFile models.PathFile

	result := db.DB.Where(models.PathFile{CubicacionID: cubicacionId}).First(&pathFile)

	if result.Error != nil {
		http.Error(w, "Error getting PathFile for put: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(p.Name) > 1 {
		pathFile.Name = p.Name
	}

	if p.Active != pathFile.Active {
		pathFile.Active = p.Active
	}

	result = db.DB.Save(&pathFile)

	if result.Error != nil {
		http.Error(w, "Error saving PathFile: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func DeletePathFilesByCubicacionId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	pathFileId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	var p models.PathFile

	result := db.DB.First(&p, pathFileId)

	if result.Error != nil {
		http.Error(w, "Error getting pathFile for deleting: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&p)

	if result.Error != nil {
		http.Error(w, "Error deleting pathFile: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
