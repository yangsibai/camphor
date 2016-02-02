package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"time"
)

var config Configuration
var ren *render.Render

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/posts", AllPosts)
	router.GET("/post", AddPostPage)
	router.POST("/post", AddPost)
	log.Fatal(http.ListenAndServe(config.Addr, router))
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ren.HTML(w, http.StatusOK, "index", "Camphor")
}

func AllPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := getAllPosts()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if posts == nil || len(posts) == 0 {
		fmt.Fprint(w, "no posts")
		return
	}

	ren.HTML(w, http.StatusOK, "posts", posts)
}

func AddPostPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ren.HTML(w, http.StatusOK, "add_post", nil)
}

func AddPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post := Post{
		ID:        bson.NewObjectId(),
		Body:      r.FormValue("body"),
		CreatedAt: time.Now(),
	}
	err := storePost(&post)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	http.Redirect(w, r, "/posts", http.StatusTemporaryRedirect)
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
