package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Body      string        `json:"body" bson:"body"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	Resources []Resource    `json:"resources" bson:"resources"`
}
