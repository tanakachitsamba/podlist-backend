package models 

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Song individual song 
	Song struct {
		Id bson.ObjectId `json:"id" bson:"_id"` 
		Title            string    `json:"title" bson:"title"`  // title of song
		Artist string `json:"artist" bson:"artist"` // artist on the song, it would be content owner or creator when its a playlist 

	}

	// Songs multiple songs, could be used for creating a playlist of songs but would require more data like ID and Content Creator 
	Songs []Song
)