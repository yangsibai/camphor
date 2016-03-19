package routes

import (
	"encoding/json"
	"fmt"
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

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := db.GetAllPosts()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	ren.HTML(w, http.StatusOK, "index", posts)
}

func AddPostPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ren.HTML(w, http.StatusOK, "add_post", nil)
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

func init() {
	ren = render.New(render.Options{
		Directory: "tmpls",
		Layout:    "layout",
	})
}
