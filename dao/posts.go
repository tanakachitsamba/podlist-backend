package dao

import (
	. "github.com/tanaka/podlist-backend/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PostsCollections : table of the jobrole
const PostsCollections = "posts"

// FindAllPosts :
func (m *DatabaseDAO) FindAllPosts() (Posts, error) {
	var posts Posts
	err := db.C(PostsCollections).Find(bson.M{}).All(&posts)
	return posts, err
}

// FindPostById : Find by its id
func (m *DatabaseDAO) FindPostById(id string) (Post, error) {
	var post Post
	err := db.C(PostsCollections).FindId(bson.ObjectIdHex(id)).One(&post)
	return post, err
}

// InsertPost : a movie into database
func (m *DatabaseDAO) InsertPost(post Post) error {
	// mongo indexing
	index := mgo.Index{
		Key:        []string{"id"}, // needs other unique identifiers as keys
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db.C(PostsCollections).EnsureIndex(index)

	err := db.C(PostsCollections).Insert(&post)

	return err
}

// DeletePost :  an existing movie
func (m *DatabaseDAO) DeletePost(post Post) error {
	err := db.C(PostsCollections).Remove(&post)
	return err
}

// UpdatePost an existing movie
func (m *DatabaseDAO) UpdatePost(post Post) error {
	err := db.C(PostsCollections).UpdateId(post.Id, &post)
	return err
}
