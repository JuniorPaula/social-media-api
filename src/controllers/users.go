package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.HandlerUser(); err != nil {
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
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connect()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	users, err := repository.Find(nameOrNickname)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, users)
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	user, err := repository.FindById(userID)

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update a user :)"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete a user :)"))
}
