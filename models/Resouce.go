package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Resource struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	BaseDir   string        `json:"-" bson: "baseDir"`
	Path      string        `json:"path" bson:"path" `
	Type      string        `json:"type" bson:"type"`
	Extension string        `json:"extension" bson:"extension"`
	Width     int           `json:"width" bson:"width"`   // only for image
	Height    int           `json:"height" bson:"height"` // only for image
	URL       string        `json:"URL" bson:"URL"`
	Hash      HashInfo      `json:"-" bson:"hash"`
	Size      int64         `json:"size" bson:"size"`
	CreatedAt time.Time     `json:"created" bson: "created"`
}
