package models

type LoginResponse struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Token    string `json:"token"`
}
