package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a user :)"))
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
