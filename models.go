package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Body      string        `json:"body" bson:"body"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

type Configuration struct {
	Addr     string `json: "addr"`
	MongoURL string `json: "mongo"`
}
