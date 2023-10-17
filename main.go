package main

import (
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/handlers"
	"github.com/rs/cors"
)

func main() {

	db.DbConnection()
	db.DbMigration()

	cors := cors.AllowAll()

	handlers.AppHandlers()

	http.ListenAndServe(":8080", cors.Handler(handlers.R))
}
