package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := getAllPosts()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if posts == nil || len(posts) == 0 {
		fmt.Fprint(w, "no posts")
		return
	}

	ren.HTML(w, http.StatusOK, "index", posts)
}

//func AllPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//posts, err := getAllPosts()
//if err != nil {
//fmt.Fprint(w, err.Error())
//return
//}
//if posts == nil || len(posts) == 0 {
//fmt.Fprint(w, "no posts")
//return
//}

//ren.HTML(w, http.StatusOK, "posts", posts)
//}

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
	ren.Text(w, http.StatusOK, "ok")
}
