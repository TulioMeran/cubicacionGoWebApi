package main

import (
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/handlers"
)

func main() {

	db.DbConnection()
	db.DbMigration()

	handlers.AppHandlers()

	http.ListenAndServe(":8080", handlers.R)
}
