package utils

import (
	"encoding/json"
	"github.com/camphor/models"
	"github.com/gorilla/sessions"
	"log"
	"os"
)

func readConfig() (error, models.Configuration) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := models.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
		return err, configuration
	}
	return nil, configuration
}

var Config models.Configuration
var store *sessions.CookieStore

func init() {
	var err error
	err, Config = readConfig()
	if err != nil {
		panic(err)
	}
	store = sessions.NewCookieStore([]byte(Config.Auth.Secret))
}
