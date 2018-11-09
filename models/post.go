package models

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/services"
	"log"
	"strings"
	"time"
)

type Post struct {
	Author string `json:"author"`
	Created time.Time `json:"created"`
	Forum string `json:"forum"`
	Id int `json:"id"`
	IsEdited bool `json:"isEdited"`
	Message string `json:"message"`
	Parent int `json:"parent"` //идентификтор родительского сообщение
	Thread int `json:"thread"`
	Path []int `json:"path"`
}

type PostDetails struct {
	Post *Post `json:"post"`
	Thread *Thread `json:"thread"`
	Forum *Forum `json:"forum"`
	Author *User `json:"author"`
}

func GetPostsIDByThreadID(db *sql.DB, threadID int) ([]int, error) {
	rows, err := db.Query("select id from posts where thread = $1", threadID)
	result := make([]int,0)
	defer rows.Close()
	if err != nil {
		funcname := services.GetFunctionName()
		log.Printf("Function: %s, Error: %v",funcname , err)
		return []int{}, err
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			funcname := services.GetFunctionName()
			log.Printf("Function: %s, Error: %v, while scaning",funcname , err)
			return []int{}, err
		}
		result = append(result,id)
	}
	return result,nil
}

func GetPathById(id int) ([]int, error) {
	db := database.GetDataBase()
	var result []int = make([]int,0)
	rows, err := db.Query(`SELECT path FROM posts WHERE id = $1`, id)
	defer rows.Close()
	if err != nil {
		funcname := services.GetFunctionName()
		log.Printf("Function: %s, Error: %v, while scaning",funcname , err)
		return []int{}, err
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			funcname := services.GetFunctionName()
			log.Printf("Function: %s, Error: %v, while scaning",funcname , err)
			return []int{}, err
		}
		result = append(result,id)
	}
	return result, nil
}



func CreatePosts(db *sql.DB,posts []*Post) ([]int, error) {
	valueStrings := make([]string, 0, len(posts))
	valueArgs := make([]interface{}, 0, len(posts) * 7)
	for i, post := range posts {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d)",i*7+1,i*7+2,i*7+3,i*7+4,i*7+5,i*7+6,i*7+7))
		valueArgs = append(valueArgs, post.Author)
		valueArgs = append(valueArgs, post.Created)
		valueArgs = append(valueArgs, post.Forum)
		valueArgs = append(valueArgs, post.Message)
		valueArgs = append(valueArgs, post.Parent)
		valueArgs = append(valueArgs, post.Thread)
		valueArgs = append(valueArgs, pq.Array(post.Path))
		fmt.Println(post.Thread,post.Forum)
	}

	stmt := fmt.Sprintf("INSERT INTO posts (author,created,forum,message,parent,thread,path) VALUES %s returning id", strings.Join(valueStrings, ","))
	fmt.Println("stmt:",stmt)
	fmt.Println("valueArgs", valueArgs)
	rows, err := db.Query(stmt,valueArgs...)
	if err != nil {
		funcname := services.GetFunctionName()
		log.Printf("Function: %s, Error: %v, while scaning",funcname , err)
		return []int{}, err
	}
	defer rows.Close()
	result := make([]int,0)
	fmt.Println("check after Query")
	fmt.Println(rows)
	for rows.Next() {
		id := 0
		err = rows.Scan(&id)
		fmt.Println("check after scan")
		if err != nil {
			funcname := services.GetFunctionName()
			log.Printf("Function: %s, Error: %v, while scaning",funcname , err)
			return []int{}, err
		}
		result = append(result,id)
	}
	fmt.Println("RESULT IDS CREATEPOSTS",result)
	if err != nil {
		return []int{},err
	}
	return result,nil
}