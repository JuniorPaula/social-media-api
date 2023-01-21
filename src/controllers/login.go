package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	userFromDB, err := repository.FindByEmail(user.Email)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if err = auth.Compare(userFromDB.Password, user.Password); err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, err)
		return
	}
	token, err := auth.TokenGenerator(userFromDB.ID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}
