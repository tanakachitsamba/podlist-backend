package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/tanaka/uphoria/dao"
	. "github.com/tanaka/uphoria/models"
)

var dao = DatabaseDAO{}

// AllSongsEndPoint :
func AllSongsEndPoint(w http.ResponseWriter, r *http.Request) {
	songs, err := dao.FindAllSongs()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, songs)
}

// FindSongEndPoint : GET a jobrole by its ID
func FindSongEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	song, err := dao.FindSongsById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobrole ID")
		return
	}
	respondWithJson(w, http.StatusOK, song)
}

// CreateSongEndPoint : POST a new jobrole
func CreateSongEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var song Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	song.Id = bson.NewObjectId()

	if err := dao.InsertSong(song); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, song)
}

// UpdateSongEndPoint : PUT update an existing jobrole
func UpdateSongEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var song Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateSong(song); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteSongEndPoint : DELETE an existing jobrole
func DeleteSongEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var song Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.DeleteSong(song); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
