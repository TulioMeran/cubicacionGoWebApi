package dto

type User struct {
	Codigo   int    `json:"codigo"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
}
