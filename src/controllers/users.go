package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.Users
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.CreateUsersRepository(db)
	repository.Create(user)
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
