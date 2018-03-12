package dao

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

// DatabaseDAO :
type DatabaseDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Connect : Establish a connection to database
func (m *DatabaseDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("Session created")

	db = session.DB(m.Database)
}
