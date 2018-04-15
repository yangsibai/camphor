package main

import (
	"github.com/camphor/routes"
	"github.com/camphor/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.Index)

	router.GET("/post", routes.AddPostPage)
	router.POST("/post", routes.HandlePost)

	router.GET("/login", routes.HandleLoginPage)
	router.POST("/login", routes.HandleLogin)

	router.GET("/logout", routes.HandleLogOut)

	router.GET("/m/:id", routes.HandleSinglePost)

	router.ServeFiles("/public/*filepath", http.Dir("public"))

	log.Println("camphor listening at", utils.Config.Addr)
	err := http.ListenAndServe(utils.Config.Addr, router)
	if err != nil {
		panic(err)
	}
}
