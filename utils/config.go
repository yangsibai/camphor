package utils

import (
	"encoding/json"
	"github.com/camphor/models"
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

func init() {
	var err error
	err, Config = readConfig()
	if err != nil {
		panic(err)
	}
}
