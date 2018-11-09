package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/models"
	"github.com/malefaro/technopark-db-forum/services"
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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
	json.Unmarshal(body,&posts)
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
	err = db.QueryRow(`SELECT MAX(id) FROM posts`).Scan(&maxId)
	maxId++
	ids, err := models.GetPostsIDByThreadID(db,thread.ID)
	fmt.Println("len posts:",len(posts))
	for i, post := range posts {
		if post.Parent != 0 && !contains(ids,post.Parent){
			t.Ctx.Output.SetStatus(http.StatusConflict)
			t.Data["json"] = &models.Error{"post parent was created in another thread"}
			t.ServeJSON()
			return
		}
		post.Thread = thread.ID
		post.Forum = thread.Forum
		post.Created = currentTime
		user, err := models.GetUserByNickname(db, post.Author)
		if err != nil {
			log.Printf("PATH: %v, error: %v", t.Ctx.Input.URI(), err)
			return
		}
		if user == nil {
			t.Data["json"] = &models.Error{"Can't find user with nickname " + post.Author}
			t.Ctx.Output.SetStatus(http.StatusNotFound)
			t.ServeJSON()
			return
		}
		parentPathes, err := models.GetPathById(post.Parent)
		post.Path = append(post.Path, parentPathes...)
		post.Path = append(post.Path, maxId+i)
		fmt.Printf("post %d: %v\n", i,post)
	}
	if len(posts) == 0 {
		post := &models.Post{}
		post.Thread = thread.ID
		post.Forum = thread.Forum
		post.Created = currentTime
		db.QueryRow(`INSERT INTO posts (forum, thread, path) VALUES($1, $2, $3) RETURNING id`, post.Forum, post.Thread, pq.Array(post.Path)).Scan(&post.Id)
	} else {
		ids, err :=models.CreatePosts(db, posts)
		if err != nil {
			funcname := services.GetFunctionName()
			log.Printf("Function: %s, Error: %v",funcname , err)
			log.Println("_____________________________________")
			
		}
		for i, id := range ids {
			posts[i].Id = id
		}
	}
	t.Ctx.Output.SetStatus(http.StatusCreated)
	t.Data["json"] = posts
	t.ServeJSON()
	
}