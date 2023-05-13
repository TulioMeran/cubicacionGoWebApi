package handlers

import (
	"github.com/TulioMeran/cubicacionGoWebApi/routes"
	"github.com/gorilla/mux"
)

var R *mux.Router

func AppHandlers() {

	R = mux.NewRouter()

	//Projects handler
	R.HandleFunc("/projects", routes.GetProjectsHandler).Methods("GET")
	R.HandleFunc("/project", routes.PostProjectHandler).Methods("POST")
	R.HandleFunc("/project", routes.DeleteProjectHandler).Methods("DELETE")
	//Status handler
	R.HandleFunc("/status", routes.GetStatusCubicacionsHandler).Methods("GET")
	R.HandleFunc("/status", routes.PostStatusCubicacionHandler).Methods("POST")
	R.HandleFunc("/status", routes.DeleteStatusCubicacion).Methods("DELETE")
	//Cubicacion handler
	R.HandleFunc("/cubicacion", routes.GetCubicacionesHandler).Methods("GET")
	R.HandleFunc("/cubicacion/file", routes.GetCubicacionFileHandler).Methods("GET")
	R.HandleFunc("/cubicacion", routes.PostCubicacionesHandler).Methods("POST")
	R.HandleFunc("/cubicacion/upload", routes.UploadCubicacionHandler).Methods("POST")
	R.HandleFunc("/cubicacion", routes.DeleteCubicacionHandler).Methods("DELETE")

}
