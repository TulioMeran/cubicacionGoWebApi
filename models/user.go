package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	LastName string `gorm:"not null" json:"lastname"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password,omitempty"`
}
