package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	Post struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Title string        `json:"title" bson:"title"`
	}

	// Users All users represented as an array of User
	Posts []Post
)
