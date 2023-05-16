package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TulioMeran/cubicacionGoWebApi/db"
	"github.com/TulioMeran/cubicacionGoWebApi/dto"
	"github.com/TulioMeran/cubicacionGoWebApi/models"
	"github.com/TulioMeran/cubicacionGoWebApi/utils"
)

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var t models.User

	json.NewDecoder(r.Body).Decode(&t)

	if len(t.Name) < 1 {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if len(t.LastName) < 1 {
		http.Error(w, "lastname is required", http.StatusBadRequest)
		return
	}

	if len(t.Email) < 1 {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 1 {
		http.Error(w, "password is required", http.StatusBadRequest)
		return
	}

	var err error
	t.Password, err = utils.PasswordEncriptor(t.Password)

	if err != nil {
		http.Error(w, "Error occurrs encripting the password in PostUserHandler: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var userEncontrado models.User

	result := db.DB.Where("email = ?", t.Email).First(&userEncontrado)

	if userEncontrado.ID != 0 {
		http.Error(w, "Email is already used for a user", http.StatusBadRequest)
		return
	}

	result = db.DB.Create(&t)

	if result.Error != nil {
		http.Error(w, "Error occurrs creating user: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	t.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(t)

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.DB.Find(&users)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting users in getuserhandler: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var usersDto []dto.User

	for _, user := range users {

		usersDto = append(usersDto, dto.User{
			Codigo:   int(user.ID),
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&usersDto)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		http.Error(w, "Error occurrs getting users for putUserHandler: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {
		http.Error(w, "User not found ", http.StatusNotFound)
		return
	}

	var t models.User
	json.NewDecoder(r.Body).Decode(&t)

	if len(t.Name) > 1 {
		user.Name = t.Name
	}

	if len(t.LastName) > 1 {
		user.LastName = t.LastName
	}

	if len(t.Email) > 1 {
		user.Email = t.Email
	}

	result = db.DB.Save(&user)

	user.Password = ""

	if result.Error != nil {
		http.Error(w, "Error occurrs saving user for putUserHandler: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.First(&user, id)

	if user.ID == 0 {
		http.Error(w, "User not found ", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error occurrs getting user for delete: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&user)

	if result.Error != nil {
		http.Error(w, "Error occurrs deleting user: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
