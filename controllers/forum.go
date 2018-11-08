package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/models"
	"net/http"
)

// Operations about Users
type ForumController struct {
	beego.Controller
}

// @Title Post
// @Description create forum
// @Param forum body models.Forum true "profile"
// @Success 201 {object} models.Forum
// @Failure 404 no such user
// @Failure 409 already exists
// @router /create [post]
func (f *ForumController) Post() {
	db := database.GetDataBase()
	body := f.Ctx.Input.RequestBody
	forum := &models.Forum{}
	json.Unmarshal(body, forum)
	user, _ := models.GetUserByNickname(db, forum.Author)
	if user == nil {
		f.Ctx.Output.SetStatus(http.StatusNotFound)
		f.Data["json"] = &models.Error{"Can't find user with nickname "+forum.Author}
		f.ServeJSON()
		return
	}
	forum.Author = user.Nickname
	err := models.CreateForum(db, forum)
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == "23505" {
			f.Ctx.Output.SetStatus(http.StatusConflict)
			forum, err := models.GetForumBySlug(db, forum.Slug)
			if err != nil { return }
			f.Data["json"] = forum
			f.ServeJSON()
			return
		}
	}
	f.Data["json"] = forum
	f.Ctx.Output.SetStatus(http.StatusCreated)
	f.ServeJSON()
}

// @Title Get
// @Description get forum
// @Param slug path string true "identificator"
// @Success 200 {object} models.Forum
// @Failure 404 no such forum
// @router /:slug/details [Get]
func (f *ForumController) Details() {
	db := database.GetDataBase()
	slug := f.GetString(":slug")
	forum , _ := models.GetForumBySlug(db, slug)
	if forum == nil {
		f.Ctx.Output.SetStatus(http.StatusNotFound)
		f.Data["json"] = &models.Error{"Can't find forum by slug "+ slug}
		f.ServeJSON()
		return
	}
	f.Ctx.Output.SetStatus(http.StatusOK)
	f.Data["json"] = forum
	f.ServeJSON()
}