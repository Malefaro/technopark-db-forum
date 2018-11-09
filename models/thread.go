package models

import (
	"database/sql"
	"fmt"
	"github.com/malefaro/technopark-db-forum/services"
	"log"
	"time"
)

type Thread struct {
	Author string `json:"author"`
	Created time.Time `json:"created"`
	Forum string `json:"forum"`
	ID int `json:"id"`
	Message string `json:"message"`
	Slug string `json:"slug"`
	Title string `json:"title"`
	Votes int `json:"votes"`
}

func (t *Thread) scanThread(rows *sql.Rows) error {
	if rows.Next() == true {
		err := rows.Scan(&t.Author,&t.Created,&t.Forum,&t.ID,&t.Message,&t.Slug,&t.Title,&t.Votes)
		if err != nil {
			log.Println("Error in scanThread:", err)
			return err
		}
		for rows.Next() {
			err := rows.Scan(&t.Author,&t.Created,&t.Forum,&t.ID,&t.Message,&t.Slug,&t.Title,&t.Votes)
			if err != nil {
				log.Println("Error in scanThread:", err)
				return err
			}
		}
	} else {
		return sql.ErrNoRows
	}
	return nil
}

func (t *Thread) scanThreads(rows *sql.Rows) error {
	err := rows.Scan(&t.Author,&t.Created,&t.Forum,&t.ID,&t.Message,&t.Slug,&t.Title,&t.Votes)
	if err != nil {
		log.Println("Error in scanThreads:",err)
		return err
	}
	return nil
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid: true,
	}
}

func CreateThread(db *sql.DB, thread *Thread) error {
	err := db.QueryRow("insert into threads (author,created,forum,message,slug,title,votes) values ($1,$2,$3,$4,$5,$6,$7) RETURNING id",
		thread.Author,thread.Created,thread.Forum,thread.Message,NewNullString(thread.Slug),thread.Title,thread.Votes).Scan(&thread.ID)
	if err != nil {
		funcname := services.GetFunctionName()
		log.Printf("Function: %s, Thread %v, Error: %v",funcname ,thread, err)
		return err
	}
	return nil
}

func GetThreadBySlug(db *sql.DB, slug string) (*Thread,error) {
	rows,err := db.Query("select * from threads where slug = $1", slug)
	defer rows.Close()
	if err != nil {
		funcname := services.GetFunctionName()
		log.Printf("Function: %s, Error: %v",funcname , err)
		return nil, err
	}
	thread := &Thread{}
	err = thread.scanThread(rows)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return thread, nil
	default:
		funcname:=services.GetFunctionName()
		log.Printf("Function: %s, Error: %v",funcname , err)
		return nil, err
	}
}

func GetThreadsByForum(db *sql.DB, forumslug,limit,since,desc string) ([]*Thread, error){
	var rows *sql.Rows
	var err error
	result := make([]*Thread, 0)
	cmp := ""
	if desc == "true" {
		desc = "DESC"
		cmp = "<="
	} else {
		desc = "ASC"
		cmp = ">="
	}
	queryrow := ""
	switch {

	case since != "" && limit != "":
		queryrow = fmt.Sprintf("select * from threads where forum = $1 and created %s $2 order by created %s limit $3", cmp,desc)
		rows, err = db.Query(queryrow, forumslug, since, limit)
	case since == "" && limit == "":
		queryrow = fmt.Sprintf("select * from threads where forum = $1 order by created %s", desc)
		rows, err = db.Query(queryrow, forumslug)
	case since == "" && limit != "":
		queryrow = fmt.Sprintf("select * from threads where forum = $1 order by created %s limit $2", desc)
		rows, err = db.Query(queryrow, forumslug, limit)
	case since != "" && limit == "":
		queryrow = fmt.Sprintf("select * from threads where forum = $1 and created %s $2 order by created %s", cmp,desc)
		rows, err = db.Query(queryrow, forumslug,since)
	}
	defer rows.Close()
	if err != nil {
		funcname:=services.GetFunctionName()
		log.Printf("Function: %s, Error: %v",funcname , err)
		return result, err
	}
	if rows.Next() == true {
		thread := &Thread{}
		err := thread.scanThreads(rows)
		if err != nil {
			return result, err
		}
		result = append(result, thread)
		for rows.Next() {
			thread := &Thread{}
			err := thread.scanThreads(rows)
			if err != nil {
				return result, err
			}
			result = append(result, thread)
		}
	} else {
		return []*Thread{}, sql.ErrNoRows
	}

	return result, nil
}