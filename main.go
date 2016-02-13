package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
)

var config Configuration
var ren *render.Render

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/posts", AllPosts)
	router.GET("/post", AddPostPage)
	router.POST("/post", AddPost)
	router.POST("/posts", AllPosts)
	log.Println("camphor run at ", config.Addr)
	log.Fatal(http.ListenAndServe(config.Addr, router))
}

func readConfig() (error, Configuration) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
		return err, configuration
	}
	return nil, configuration
}

func init() {
	var err error
	err, config = readConfig()
	if err != nil {
		log.Fatal(err)
	}
	ren = render.New(render.Options{
		Directory: "tmpls",
	})
}
