package dto

type Cubicacion struct {
	Codigo           int    `json:"codigo"`
	Descripcion      string `json:"descripcion"`
	Observacion      string `json:"observacion"`
	Ruta             string `json:"ruta"`
	Proyecto         Project
	EstadoCubicacion StatusCubicacion
	Comments         []Comment
}
