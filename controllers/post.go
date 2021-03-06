package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/models"
	"net/http"
	"strings"
)

type PostController struct {
	beego.Controller
}

// @Title Post
// @Description create forum
// @Param id path string true "id"
// @Param related query bool false "related"
// @Success 201 {object} models.Forum
// @Failure 404 no such user
// @Failure 409 already exists
// @router /:id/details [get]
func (p *PostController) Get() {
	db := database.GetDataBase()

	id := p.GetString(":id")
	related := p.Ctx.Input.Query("related")
	infos := strings.Split(related, ",")
	//fmt.Println("INFOS in controller",infos)
	//pd,err := models.GetPostDetailsByID(db, id)
	//if err != nil && err != sql.ErrNoRows {
	//	return
	//}
	//if pd == nil {
	//	p.Ctx.Output.SetStatus(http.StatusNotFound)
	//	p.Data["json"] = &models.Error{"Can't fild post with id: "+id}
	//	p.ServeJSON()
	//	return
	//}

	basequery := `select * from posts as p `
	addUser := ""
	addForum := ""
	addThread := ""
	result := &models.PostDetails{}
	for _, v := range infos {
		if v == "user" {
			//result.Author = pd.Author
			addUser = " join users as u on u.nickname = p.author "
			basequery += addUser
		}
		if v == "forum" {
			//result.Forum = pd.Forum
			addForum = " join forums as f on p.forum = f.slug "
			basequery += addForum
		}
		if v == "thread" {
			//result.Thread = pd.Thread
			addThread = " join threads thread on p.thread = thread.id "
			basequery += addThread
		}
	}
	basequery += "where p.id = $1"
	//querystr := fmt.Sprintf(basequery, addUser, addForum, addThread)
	querystr := basequery
	//fmt.Println(basequery)
	pd, err  := models.GetPostDetailsByIDrework(db, querystr, infos, id)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if pd == nil {
		p.Ctx.Output.SetStatus(http.StatusNotFound)
		p.Data["json"] = &models.Error{"Can't fild post with id: "+id}
		p.ServeJSON()
		return
	}
	for _, v := range infos {
		if v == "user" {
			result.Author = pd.Author
		}
		if v == "forum" {
			result.Forum = pd.Forum
		}
		if v == "thread" {
			result.Thread = pd.Thread
		}
	}
	result.Post = pd.Post
	p.Ctx.Output.SetStatus(http.StatusOK)
	//p.Data["json"] = result
	//p.ServeJSON()
	serveJson(result, p.Ctx.Output)
}

// @Title Post
// @Description create forum
// @Param id path string true "id"
// @Param related query bool false "related"
// @Success 201 {object} models.Forum
// @Failure 404 no such user
// @Failure 409 already exists
// @router /:id/details [post]
func (p *PostController) UpdatePosts() {
	db := database.GetDataBase()
	id := p.GetString(":id")
	body := p.Ctx.Input.RequestBody
	updatepost := &models.Post{}
	json.Unmarshal(body, updatepost)
	//easyjson.Unmarshal(body, updatepost)
	//tx, err := db.Begin()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer tx.Commit()
	//pathstring:=""
	//post := &models.Post{}
	//err = tx.QueryRow("select * from posts where id = $1", id).Scan(&post.Author,&post.Created,&post.Forum,&post.Id,&post.IsEdited,&post.Message,&post.Parent,&post.Thread,&pathstring)
	//if err != nil {
	//	p.Ctx.Output.SetStatus(http.StatusNotFound)
	//	p.Data["json"] = &models.Error{"can't find post with id: "+ id}
	//	p.ServeJSON()
	//	return
	//}
	//if updatepost.Message != "" && post.Message != updatepost.Message {
	//	post, err = models.UpdatePosts(db, post.Message, id)
	//	if err != nil {
	//		p.Ctx.Output.SetStatus(http.StatusNotFound)
	//		p.Data["json"] = &models.Error{"can't find post with id: "+ id}
	//		p.ServeJSON()
	//		return
	//	}
	//}
	post, err := models.UpdatePosts(db, updatepost.Message, id)
	if post == nil {
		p.Ctx.Output.SetStatus(http.StatusNotFound)
		p.Data["json"] = &models.Error{"can't find post with id: " + id}
		p.ServeJSON()
		return
	}
	if err != nil {
		p.Ctx.Output.SetStatus(http.StatusNotFound)
		p.Data["json"] = &models.Error{"can't find post with id: " + id}
		p.ServeJSON()
		return
	}
	p.Ctx.Output.SetStatus(http.StatusOK)
	//p.Data["json"] = post
	//p.ServeJSON()
	serveJson(post, p.Ctx.Output)
}