package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
}
