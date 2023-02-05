package controllers

import (
	"api-dvbk-socialNetwork/src/auth"
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/models"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// --
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userTokenId, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	var post models.Post
	if err := json.Unmarshal(bodyRequest, &post); err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.AuthorID = userTokenId

	if err := post.PreparePostData(); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewPostsRepository(DB)
	post.ID, err = repository.CreatePost(post)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 201, post)

}

// --
func GetPosts(w http.ResponseWriter, r *http.Request) {

}

// --
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewPostsRepository(DB)
	post, err := repository.SearchPost(postID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, post)
}

// --
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// --
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
