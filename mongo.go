package main

import (
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(config.MongoURL)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func storePost(post *Post) error {
	session := getSession()
	defer session.Close()

	C := session.DB("camphor").C("post")
	err := C.Insert(&post)
	return err
}

func getAllPosts() (posts []Post, err error) {
	session := getSession()
	defer session.Close()

	C := session.DB("camphor").C("post")
	err = C.Find(nil).Sort("-created_at").All(&posts)
	return
}
