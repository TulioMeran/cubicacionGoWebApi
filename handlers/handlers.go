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
	R.HandleFunc("/project", routes.PutProjectHandler).Methods("PUT")
	R.HandleFunc("/project", routes.DeleteProjectHandler).Methods("DELETE")
	//Status handler
	R.HandleFunc("/status", routes.GetStatusCubicacionsHandler).Methods("GET")
	R.HandleFunc("/status", routes.PostStatusCubicacionHandler).Methods("POST")
	R.HandleFunc("/status", routes.PutStatusCubicacion).Methods("PUT")
	R.HandleFunc("/status", routes.DeleteStatusCubicacion).Methods("DELETE")
	//Cubicacion handler
	R.HandleFunc("/cubicacion", routes.GetCubicacionesHandler).Methods("GET")
	R.HandleFunc("/cubicacion/file", routes.GetCubicacionFileHandler).Methods("GET")
	R.HandleFunc("/cubicacion", routes.PostCubicacionesHandler).Methods("POST")
	R.HandleFunc("/cubicacion/upload", routes.UploadCubicacionHandler).Methods("POST")
	R.HandleFunc("/cubicacion", routes.DeleteCubicacionHandler).Methods("DELETE")
	R.HandleFunc("/cubicacion", routes.PutCubicacionHandler).Methods("PUT")

	//User handler
	R.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	R.HandleFunc("/user", routes.PutUserHandler).Methods("PUT")
	R.HandleFunc("/user", routes.GetUsersHandler).Methods("GET")
	R.HandleFunc("/user", routes.DeleteUserHandler).Methods("DELETE")

	//Auth handler
	R.HandleFunc("/auth/login", routes.LoginHandler).Methods("POST")

}
