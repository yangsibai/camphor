package routes

import (
	"encoding/json"
	"github.com/camphor/db"
	"github.com/camphor/models"
	"github.com/camphor/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

var ren *render.Render

func init() {
	ren = render.New(render.Options{
		Directory: "tmpls",
		Layout:    "layout",
	})
}

func responseError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	posts, err := db.GetAllPosts()
	if err != nil {
		responseError(w, err)
		return
	}
	isLogin, err := utils.IsLogin(r)
	if err != nil {
		responseError(w, err)
		return
	}

	loc, err := time.LoadLocation("Asia/ShangHai")
	if err != nil {
		responseError(w, err)
		return
	}
	for i := 0; i < len(posts); i++ {
		posts[i].CreatedAt = posts[i].CreatedAt.In(loc)
	}
	ren.HTML(w, http.StatusOK, "index", struct {
		Posts   []models.Post
		IsLogin bool
	}{
		posts,
		isLogin,
	})
}

func AddPostPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authed, err := utils.IsLogin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if authed == true {
		ren.HTML(w, http.StatusOK, "add_post", nil)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := []byte(r.PostFormValue("data"))

	var post models.Post
	err := json.Unmarshal(data, &post)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	post.ID = bson.NewObjectId()
	post.CreatedAt = time.Now()

	if len(post.Resources) > 0 {
		for i := 0; i < len(post.Resources); i++ {
			post.Resources[i].ID = bson.NewObjectId()
			post.Resources[i].CreatedAt = time.Now()
		}
	}

	err = db.StorePost(&post)
	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}
	utils.WriteResponse(w, post.ID.Hex())
}

func HandleLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ren.HTML(w, http.StatusOK, "login", nil)
}

func HandleLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	success, err := utils.Login(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if success {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ren.HTML(w, http.StatusOK, "login", "login failed")
}
