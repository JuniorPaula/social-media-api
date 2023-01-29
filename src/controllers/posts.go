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
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Validations(); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, post)
}

func FindAllPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.FindAll(userID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, posts)

}

func FindPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostsRepository(db)
	post, err := repository.FindById(postID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {}

func DeletePost(w http.ResponseWriter, r *http.Request) {}
