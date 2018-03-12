package dao

import (
	. "github.com/tanaka/uphoria/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SongsCollection : table of the jobrole
const SongsCollection = "songs"

// FindAllSongs :
func (m *DatabaseDAO) FindAllSongs() (Songs, error) {
	var songs Songs
	err := db.C(SongsCollection).Find(bson.M{}).All(&songs)
	return songs, err
}

// FindSongsById : Find by its id
func (m *DatabaseDAO) FindSongsById(id string) (Song, error) {
	var song Song
	err := db.C(SongsCollection).FindId(bson.ObjectIdHex(id)).One(&song)
	return song, err
}

// InsertSong : a movie into database
func (m *DatabaseDAO) InsertSong(song Song) error {
	// mongo indexing
	index := mgo.Index{
		Key:        []string{"id"}, // needs other unique identifiers as keys
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db.C(SongsCollection).EnsureIndex(index)

	err := db.C(SongsCollection).Insert(&song)

	return err
}

// DeleteSong :  an existing movie
func (m *DatabaseDAO) DeleteSong(song Song) error {
	err := db.C(SongsCollection).Remove(&song)
	return err
}

// UpdateSong an existing movie
func (m *DatabaseDAO) UpdateSong(song Song) error {
	err := db.C(SongsCollection).UpdateId(song.Id, &song)
	return err
}
