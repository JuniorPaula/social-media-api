package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Users
	if err = json.Unmarshal(requestBody, &user); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find a users :)"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find a user by id :)"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update a user :)"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete a user :)"))
}
