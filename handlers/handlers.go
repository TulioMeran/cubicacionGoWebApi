package handlers

import (
	"github.com/TulioMeran/cubicacionGoWebApi/middlewares"
	"github.com/TulioMeran/cubicacionGoWebApi/routes"
	"github.com/gorilla/mux"
)

var R *mux.Router

func AppHandlers() {

	R = mux.NewRouter()

	//Projects handler
	R.HandleFunc("/projects", middlewares.TokenValidator(routes.GetProjectsHandler)).Methods("GET")
	R.HandleFunc("/project", middlewares.TokenValidator(routes.PostProjectHandler)).Methods("POST")
	R.HandleFunc("/project", middlewares.TokenValidator(routes.PutProjectHandler)).Methods("PUT")
	R.HandleFunc("/project", middlewares.TokenValidator(routes.DeleteProjectHandler)).Methods("DELETE")
	//Status handler
	R.HandleFunc("/status", middlewares.TokenValidator(routes.GetStatusCubicacionsHandler)).Methods("GET")
	R.HandleFunc("/status", middlewares.TokenValidator(routes.PostStatusCubicacionHandler)).Methods("POST")
	R.HandleFunc("/status", middlewares.TokenValidator(routes.PutStatusCubicacion)).Methods("PUT")
	R.HandleFunc("/status", middlewares.TokenValidator(routes.DeleteStatusCubicacion)).Methods("DELETE")
	//Cubicacion handler
	R.HandleFunc("/cubicacion", middlewares.TokenValidator(routes.GetCubicacionesHandler)).Methods("GET")
	R.HandleFunc("/cubicacion/file", routes.GetCubicacionFileHandler).Methods("GET")
	R.HandleFunc("/cubicacion", middlewares.TokenValidator(routes.PostCubicacionesHandler)).Methods("POST")
	R.HandleFunc("/cubicacion/upload", middlewares.TokenValidator(routes.UploadCubicacionHandler)).Methods("POST")
	R.HandleFunc("/cubicacion", middlewares.TokenValidator(routes.DeleteCubicacionHandler)).Methods("DELETE")
	R.HandleFunc("/cubicacion", middlewares.TokenValidator(routes.PutCubicacionHandler)).Methods("PUT")

	//User handler
	R.HandleFunc("/user", middlewares.TokenValidator(routes.PostUserHandler)).Methods("POST")
	R.HandleFunc("/user", middlewares.TokenValidator(routes.PutUserHandler)).Methods("PUT")
	R.HandleFunc("/user", middlewares.TokenValidator(routes.GetUsersHandler)).Methods("GET")
	R.HandleFunc("/user", middlewares.TokenValidator(routes.DeleteUserHandler)).Methods("DELETE")

	//Auth handler
	R.HandleFunc("/auth/login", routes.LoginHandler).Methods("POST")

	//Comment handler
	R.HandleFunc("/comment", middlewares.TokenValidator(routes.GetCommentsHandler)).Methods("GET")
	R.HandleFunc("/comment", middlewares.TokenValidator(routes.PostCommentHandler)).Methods("POST")
	R.HandleFunc("/comment", middlewares.TokenValidator(routes.PutCommentHandler)).Methods("PUT")
	R.HandleFunc("/comment", middlewares.TokenValidator(routes.DeleteCommentHandler)).Methods("DELETE")

	//PathFiles
	R.HandleFunc("/pathfile", middlewares.TokenValidator(routes.GetPathFilesByCubicationId)).Methods("GET")
	R.HandleFunc("/pathfile", middlewares.TokenValidator(routes.PostPathFilesByCubicacionId)).Methods("POST")
	R.HandleFunc("/pathfile", middlewares.TokenValidator(routes.PutPathFilesByCubicationId)).Methods("PUT")
	R.HandleFunc("/pathfile", middlewares.TokenValidator(routes.DeletePathFilesByCubicacionId)).Methods("DELETE")

}
