package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/malefaro/technopark-db-forum/models"
)

// custom controller
type ThreadController struct {
	beego.Controller
	DB *sql.DB
}

// @Title GetAll
// @Description get Thread from url
// @Success 200 {object} models.Thread
// @router /:slug_or_id/create [get]
func (t *ThreadController) GetAll() {
	slug_or_id := t.GetString(":slug_or_id")
	thread := models.Thread{}
	thread.AddThread(0, slug_or_id)
	t.Data["json"] = thread
	t.ServeJSON()
}

// @Title GetAll
// @Description get Thread from url
// @Success 200 {object} models.Thread
// @router /create [get]
func (t *ThreadController) Create() {
	thread := models.Thread{}
	thread.AddThread(1, "kek")
	t.Data["json"] = thread
	t.ServeJSON()
}

type Test struct {
	beego.Controller
}

func (t *Test) Test() {
	thread := models.Thread{}
	thread.AddThread(1, "kek")
	t.Data["json"] = thread
	t.ServeJSON()
}
