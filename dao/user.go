package dao

import (
	. "github.com/tanaka/uphoria/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserCollection : table of the jobrole
const UserCollection = "users"

// FindAllUsers :
func (m *DatabaseDAO) FindAllUsers() (Users, error) {
	var users Users
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	return users, err
}

// FindUsersById : Find by its id
func (m *DatabaseDAO) FindUsersById(id string) (User, error) {
	var user User
	err := db.C(UserCollection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// InsertUser : a movie into database
func (m *DatabaseDAO) InsertUser(user User) error {
	// mongo indexing
	index := mgo.Index{
		Key:        []string{"id"}, // needs other unique identifiers as keys
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db.C(UserCollection).EnsureIndex(index)

	err := db.C(UserCollection).Insert(&user)

	return err
}

// DeleteUsers :  an existing movie
func (m *DatabaseDAO) DeleteUsers(user User) error {
	err := db.C(UserCollection).Remove(&user)
	return err
}

// UpdateUsers an existing movie
func (m *DatabaseDAO) UpdateUsers(user User) error {
	err := db.C(UserCollection).UpdateId(user.Id, &user)
	return err
}
