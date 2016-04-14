package models

import (
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"time"
)

type Post struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Body      string        `json:"body" bson:"body"`
	HTML      template.HTML `json:"-" bson:"-"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	Resources []Resource    `json:"resources" bson:"resources"`
}
