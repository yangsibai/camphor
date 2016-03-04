package routes

import (
	"fmt"
	"github.com/camphor/db"
	"github.com/camphor/models"
	"github.com/camphor/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var ren *render.Render

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := db.GetAllPosts()
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

// save a single image
func handleSaveSingleResource(part *multipart.Part) (info models.Resource, err error) {
	newID := bson.NewObjectId()
	date := time.Now().Format("20060102")

	fmt.Println(part.FileName())
	err = utils.CreateDirIfNotExists(filepath.Join(utils.Config.SaveDir, date))
	if err != nil {
		return
	}
	path := filepath.Join(date, newID.Hex())
	savePath := filepath.Join(utils.Config.SaveDir, path)

	dst, err := os.Create(savePath)

	if err != nil {
		return
	}

	defer dst.Close()

	var bytes int64
	if bytes, err = io.Copy(dst, part); err != nil {
		return
	}

	ext := filepath.Ext(part.FileName())

	var width, height int

	if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
		width, height = utils.GetImageDimensions(savePath)
	}

	URL := utils.Config.BaseURL + path

	var hash models.HashInfo

	hash, err = utils.CalculateBasicHashes(savePath)

	if err != nil {
		return
	}

	info = models.Resource{
		ID:        newID,
		Name:      part.FileName(),
		Extension: ext,
		BaseDir:   utils.Config.SaveDir,
		Path:      path,
		Width:     width,
		Height:    height,
		URL:       URL,
		Hash:      hash,
		Size:      bytes,
		CreatedAt: time.Now(),
	}
	err = db.StoreResource(&info)
	if err != nil {
		return
	}
	return info, nil
}

func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	content := r.PostFormValue("content")

	post := models.Post{
		ID:        bson.NewObjectId(),
		Body:      content,
		CreatedAt: time.Now(),
	}

	err := db.StorePost(&post)
	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}
	utils.WriteResponse(w, content)
}

func AddPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//for _, fheaders := range r.MultipartForm.File {
	//for _, hdr := range fheaders {
	//if infile, err := hdr.Open(); err != nil {
	//utils.WriteErrorResponse(w, err)
	//return
	//}
	//}
	//}

	//reader, err := r.MultipartReader()
	//if err != nil {
	//utils.WriteErrorResponse(w, err)
	//return
	//}

	//var resources []models.Resource

	//for {
	//part, err := reader.NextPart()
	//if err == io.EOF {
	//break
	//}
	//if part.FileName() == "" {
	//continue
	//}
	//resource, err := handleSaveSingleResource(part)
	//resources = append(resources, resource)
	//}

	//fmt.Println(r.FormValue("body"))
	//post := models.Post{
	//ID:        bson.NewObjectId(),
	//Body:      r.FormValue("body"),
	//CreatedAt: time.Now(),
	//Resources: resources,
	//}

	//err = db.StorePost(&post)
	//if err != nil {
	//utils.WriteErrorResponse(w, err)
	//return
	//}
	ren.Text(w, http.StatusOK, "ok: "+r.PostFormValue("body"))
}

func init() {
	ren = render.New(render.Options{
		Directory: "tmpls",
		Layout:    "layout",
	})
}
