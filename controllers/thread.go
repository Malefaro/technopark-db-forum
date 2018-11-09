package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

// custom controller
type ThreadController struct {
	beego.Controller
}

// @Title GetAll
// @Description get Thread from url
// @Success 200 {object} models.Thread
// @router /:slug_or_id/details [post]
func (t *ThreadController) UpdateThread() {
	db := database.GetDataBase()
	slug_or_id := t.GetString(":slug_or_id")
	body := t.Ctx.Input.RequestBody
	thread := &models.Thread{}
	oldthread := &models.Thread{}
	json.Unmarshal(body,thread)
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		//thread.ID = id
		oldthread, err = models.GetTreadByID(db, id)
		if oldthread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with id: "+strconv.Itoa(id)}
			t.ServeJSON()
			return
		}

	} else {
		//thread.Slug = slug_or_id
		oldthread, err = models.GetThreadBySlug(db,slug_or_id)
		if oldthread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with slug: "+slug_or_id}
			t.ServeJSON()
			return
		}
	}
	if thread.Title != "" {
		oldthread.Title = thread.Title
	}
	if thread.Message != "" {
		oldthread.Message = thread.Message
	}
	err = models.UpdateThread(db, oldthread)
	if err != nil {
		return
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = oldthread
	t.ServeJSON()
}

// @Title GetThread by slug or id
// @Description get Thread from url
// @Success 200 {object} models.Thread
// @router /:slug_or_id/details [get]
func (t *ThreadController) GetThread() {
	db := database.GetDataBase()
	slug_or_id := t.GetString(":slug_or_id")
	thread := &models.Thread{}
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		//thread.ID = id
		thread, err = models.GetTreadByID(db, id)
		if thread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with id: "+strconv.Itoa(id)}
			t.ServeJSON()
			return
		}

	} else {
		//thread.Slug = slug_or_id
		thread, err = models.GetThreadBySlug(db,slug_or_id)
		if thread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with slug: "+slug_or_id}
			t.ServeJSON()
			return
		}
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = thread
	t.ServeJSON()
}

// @Title GetThread by slug or id
// @Description get Thread from url
// @Success 200 {object} models.Thread
// @router /:slug_or_id/vote [post]
func (t *ThreadController) CreateVote() {
	db := database.GetDataBase()
	slug_or_id := t.GetString(":slug_or_id")
	vote := &models.Vote{}
	body := t.Ctx.Input.RequestBody
	json.Unmarshal(body,vote)
	thread := &models.Thread{}
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		//thread.ID = id
		thread, err = models.GetTreadByID(db, id)
		if thread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with id: "+strconv.Itoa(id)}
			t.ServeJSON()
			return
		}

	} else {
		//thread.Slug = slug_or_id
		thread, err = models.GetThreadBySlug(db,slug_or_id)
		if thread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with slug: "+slug_or_id}
			t.ServeJSON()
			return
		}
	}
	vote.Thread = thread.ID
	//fmt.Println("___________________________________")
	//fmt.Println(vote)
	//fmt.Println("vote voice", vote.Voice)
	//fmt.Println("___________________________________")
	err = models.CreateVote(db, vote)
	if pgerr, ok := err.(*pq.Error); ok {
		//fmt.Printf("%v\n", pgerr)
		//fmt.Printf("%#v\n", pgerr.Code)
		if pgerr.Code == "23505" {
			voice,_ := models.UpdateVote(db, vote)
			if voice != 0{
				thread.Votes += 2*vote.Voice
			}
			t.Ctx.Output.SetStatus(http.StatusOK)
			t.Data["json"] = thread
			t.ServeJSON()
			return
		}
	}
	thread.Votes += vote.Voice
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = thread
	t.ServeJSON()
}


// @Title GetThread by slug or id
// @Description get Thread from url
// @Success 200 {object} models.Thread
// @router /:slug_or_id/create [post]
func (t *ThreadController) CreatePosts() {
	currentTime := time.Now()
	db := database.GetDataBase()
	body := t.Ctx.Input.RequestBody
	slug_or_id := t.GetString(":slug_or_id")
	posts := make([]models.Post,0)
	json.Unmarshal(body,posts)
	id, err := strconv.Atoi(slug_or_id)
	thread := &models.Thread{}
	if err == nil {
		//thread.ID = id
		thread, err = models.GetTreadByID(db, id)
		if thread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with id: "+strconv.Itoa(id)}
			t.ServeJSON()
			return
		}
	} else {
		//thread.Slug = slug_or_id
		thread, err = models.GetThreadBySlug(db,slug_or_id)
		if thread == nil {
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.Data["json"] = &models.Error{"Can't find thread with slug: "+slug_or_id}
			t.ServeJSON()
			return
		}
	}
	maxId:= 0
	maxId++
	for i := range posts {
		posts[i].Thread = thread.ID
		posts[i].Forum = thread.Forum
		posts[i].Created = currentTime
		user, err := models.GetUserByNickname(db, posts[i].Author)
		if err != nil {
			log.Printf("PATH: %v, error: %v", t.Ctx.Input.URI(), err)
			return
		}
		if user == nil {
			t.Data["json"] = &models.Error{"Can't find user with nickname " + nickname}
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.ServeJSON()
			return
		}

	}
}