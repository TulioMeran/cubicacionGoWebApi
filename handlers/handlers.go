package handlers

import (
	"github.com/TulioMeran/cubicacionGoWebApi/routes"
	"github.com/gorilla/mux"
)

var R *mux.Router

func AppHandlers() {

	R = mux.NewRouter()

	R.HandleFunc("/projects", routes.GetProjectsHandler).Methods("GET")
	R.HandleFunc("/project", routes.PostProjectHandler).Methods("POST")
	R.HandleFunc("/project", routes.DeleteProjectHandler).Methods("DELETE")
	R.HandleFunc("/status", routes.GetProjectsHandler).Methods("GET")
	R.HandleFunc("/status", routes.PostProjectHandler).Methods("POST")
	R.HandleFunc("/status", routes.DeleteProjectHandler).Methods("DELETE")

}
