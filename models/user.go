package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id          bson.ObjectId `json:"id" bson:"_id"`
		FirstName string        `json:"first_name" bson:"first_name"`
		LastName string        `json:"last_name" bson:"last_name"`
		Email       string        `json:"email bson:"email"`
	}

	// Users All users represented as an array of User
	Users []User
)
