package models

import "gorm.io/gorm"

type Cubicacion struct {
	gorm.Model

	Description        string `gorm:"not null" json:"description"`
	Observation        string `json:"observation"`
	PathFile           string `json:"pathfile"`
	ProjectID          int    `gorm:"not null" json:"projectid"`
	Project            Project
	StatusCubicacionID int `gorm:"not null" json:"statuscubicacionid"`
	StatusCubicacion   StatusCubicacion
}
