package models

import "gorm.io/gorm"

type StatusCubicacion struct {
	gorm.Model

	Description string `gorm:"not null" json:"description"`
}
