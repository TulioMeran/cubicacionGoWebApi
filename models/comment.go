package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	Description        string `gorm:"not null" json:"description"`
	UserID             int    `gorm:"not null" json:"userid"`
	CubicacionID       int    `gorm:"not null" json:"cubicacionid"`
	StatusCubicacionID int    `gorm:"not null" json:"statuscubicacionid"`
}
