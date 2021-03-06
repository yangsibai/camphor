package db

import (
	"github.com/camphor/models"
	"github.com/camphor/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(utils.Config.MongoURL)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func StorePost(post *models.Post) error {
	session := getSession()
	defer session.Close()

	C := session.DB("camphor").C("post")
	err := C.Insert(&post)
	return err
}

func GetAllPosts() (posts []models.Post, err error) {
	session := getSession()
	defer session.Close()

	C := session.DB("camphor").C("post")
	err = C.Find(nil).Sort("-created_at").All(&posts)
	return
}

func StoreResource(resource *models.Resource) error {
	session := getSession()
	defer session.Close()

	C := session.DB("camphor").C("resource")
	return C.Insert(&resource)
}

func GetSinglePost(id string) (post models.Post, err error) {
	session := getSession()

	defer session.Close()

	oid := bson.ObjectIdHex(id)
	C := session.DB("camphor").C("post")
	err = C.FindId(oid).One(&post)
	return
}

func GetPostByIndex(idx int) (post models.Post, err error) {
	session := getSession()

	defer session.Close()

	C := session.DB("camphor").C("post")
	err = C.Find(bson.M{}).Skip(idx).One(&post)
	return
}

func GetPostsCount() (count int, err error) {
	session := getSession()

	defer session.Close()

	C := session.DB("camphor").C("post")
	return C.Find(bson.M{}).Count()
}
