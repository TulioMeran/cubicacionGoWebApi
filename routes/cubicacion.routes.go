package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/dto"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func GetCubicacionesHandler(w http.ResponseWriter, r *http.Request) {
	var cubicaciones []models.Cubicacion

	result := db.DB.Preload("Project").Preload("StatusCubicacion").Preload("Comments").Preload("PathFiles").Find(&cubicaciones)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting cubicaciones: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var cubicacionesDto []dto.Cubicacion

	for _, c := range cubicaciones {
		var proyecto = dto.Project{
			Codigo:      int(c.Project.ID),
			Titulo:      c.Project.Title,
			Descripcion: c.Project.Description,
		}
		var estado = dto.StatusCubicacion{
			Codigo:      int(c.StatusCubicacion.ID),
			Descripcion: c.StatusCubicacion.Description,
		}

		var commentsDto []dto.Comment

		for _, comment := range c.Comments {
			var t = dto.Comment{
				Codigo:      int(comment.ID),
				Descripcion: comment.Description,
			}
			commentsDto = append(commentsDto, t)
		}

		var pathFilesDto []dto.Ruta

		for _, path := range c.PathFiles {
			var t = dto.Ruta{
				Nombre:        path.Name,
				Activo:        path.Active,
				FechaRegistro: path.CreatedAt,
			}
			pathFilesDto = append(pathFilesDto, t)
		}

		var element = dto.Cubicacion{
			Codigo:           int(c.ID),
			Descripcion:      c.Description,
			Observacion:      c.Observation,
			Rutas:            pathFilesDto,
			Proyecto:         proyecto,
			EstadoCubicacion: estado,
			Comments:         commentsDto,
		}

		cubicacionesDto = append(cubicacionesDto, element)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&cubicacionesDto)
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

func UploadCubicacionHandler(w http.ResponseWriter, r *http.Request) {
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

	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/cubicacion/" + id + "." + extension

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

	//	t.PathFile = id + "." + extension

	result = db.DB.Save(&t)

	//result = db.DB.Model(&t).Update("PathFile", id+"."+extension)

	if result.Error != nil {
		http.Error(w, "Error occurrs updating cubicacion in upload file: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetCubicacionFileHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("id")) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")

	var t models.Cubicacion

	result := db.DB.First(&t, id)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting cubicacion in cubicacion file: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var pathFileName string

	for _, path := range t.PathFiles {

		if path.Active {
			pathFileName = path.Name
			break
		}
	}

	openFile, err := os.Open("uploads/cubicacion/" + pathFileName)

	if err != nil {
		http.Error(w, "Error occurrs opening file in cubicacion file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error occurrs copying file in cubicacion file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutCubicacionHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var cubicacion models.Cubicacion
	var t models.Cubicacion
	json.NewDecoder(r.Body).Decode(&t)

	result := db.DB.First(&cubicacion, id)

	if cubicacion.ID < 1 {
		http.Error(w, "Cubicacion not found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting cubicacion for putcubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(t.Description) > 1 {
		cubicacion.Description = t.Description
	}

	if t.StatusCubicacionID > 0 {
		cubicacion.StatusCubicacionID = t.StatusCubicacionID
	}

	if len(t.Observation) > 1 {
		cubicacion.Observation = t.Observation
	}

	if t.ProjectID > 0 {
		cubicacion.ProjectID = t.ProjectID
	}

	result = db.DB.Save(&cubicacion)

	if result.Error != nil {
		http.Error(w, "Error occurrs saving cubicacion putCubicacion: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
