package models

type Configuration struct {
	Addr     string `json: "addr"`
	MongoURL string `json: "mongo"`
}
