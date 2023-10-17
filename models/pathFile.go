package models

import "gorm.io/gorm"

type PathFile struct {
	gorm.Model

	Name         string `gorm:"not null" json:"name"`
	Active       bool   `gorm:"not null" json:"active"`
	CubicacionID int    `gorm:"not null" json:"cubicacionid"`
}
