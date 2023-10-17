package dto

type Cubicacion struct {
	Codigo           int    `json:"codigo"`
	Descripcion      string `json:"descripcion"`
	Observacion      string `json:"observacion"`
	Rutas            []Ruta
	Proyecto         Project
	EstadoCubicacion StatusCubicacion
	Comments         []Comment
}
