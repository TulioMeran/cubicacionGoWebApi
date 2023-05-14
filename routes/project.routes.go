package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/dto"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	result := db.DB.Find(&projects)

	if result.Error != nil {
		http.Error(w, "Error happend getting projects: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var projectsDto []dto.Project

	for _, project := range projects {

		var element = dto.Project{
			Codigo:      int(project.ID),
			Titulo:      project.Title,
			Descripcion: project.Description,
		}

		projectsDto = append(projectsDto, element)
	}

	json.NewEncoder(w).Encode(&projectsDto)

}

func PostProjectHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Project
	json.NewDecoder(r.Body).Decode(&t)

	if len(t.Title) < 1 {
		http.Error(w, "title is required.", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&t)

	if result.Error != nil {
		http.Error(w, "Error occurs creating project: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&t)

}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if len(r.URL.Query().Get("id")) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")

	result := db.DB.First(&project, id)

	if result.Error != nil {
		http.Error(w, "Error occurs getting project for deleted: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.DB.Unscoped().Delete(&project)

	if result.Error != nil {
		http.Error(w, "Error occurs deleting project: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func PutProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.Project
	json.NewDecoder(r.Body).Decode(&t)

	var project models.Project

	result := db.DB.First(&project, id)

	if project.ID < 1 {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting project for put project: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(t.Title) > 1 {
		project.Title = t.Title
	}

	if len(t.Description) > 1 {
		project.Description = t.Description
	}

	result = db.DB.Save(&project)

	if result.Error != nil {
		http.Error(w, "Error occurrs saving project for put project: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
