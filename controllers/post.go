package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/tanaka/podlist-backend/dao"
	. "github.com/tanaka/podlist-backend/models"
)

var dao = DatabaseDAO{}

// AllPostsEndPoint :
func AllPostsEndPoint(w http.ResponseWriter, r *http.Request) {
	posts, err := dao.FindAllPosts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, posts)
}

// FindPostEndPoints : GET a jobrole by its ID
func FindPostEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post, err := dao.FindPostById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobrole ID")
		return
	}
	respondWithJson(w, http.StatusOK, post)
}

// CreatePostEndPoint : POST a new jobrole
func CreatePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	post.Id = bson.NewObjectId()

	if err := dao.InsertPost(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, post)
}

// UpdatePostEndPoint : PUT update an existing jobrole
func UpdatePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdatePost(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeletePostEndPoint : DELETE an existing jobrole
func DeletePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.DeletePost(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
