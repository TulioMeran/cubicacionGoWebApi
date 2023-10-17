package dto

import (
	"time"
)

type Ruta struct {
	Nombre        string    `json:"nombre"`
	Activo        bool      `json:"activo"`
	FechaRegistro time.Time `json:"fechaRegistro"`
}
