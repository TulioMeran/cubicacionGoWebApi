package db

import (
	"github.com/TulioMeran/cubicacionGoWebApi/models"
)

func DbMigration() {
	DB.AutoMigrate(&models.Comment{})
	DB.AutoMigrate(&models.Project{}, &models.StatusCubicacion{}, &models.Cubicacion{}, &models.User{})
}
