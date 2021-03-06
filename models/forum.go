package models

import (
	"database/sql"
)

//easyjson:json
type Forum struct {
	Posts int `json:"posts"`
	Slug string `json:"slug"`
	Threads int `json:"threads"`
	Title string `json:"title"`
	Author string `json:"user"`
}

var StmtGetForumBySlug *sql.Stmt
var StmtGetThreadByID *sql.Stmt
var StmtGetThreadBySlug *sql.Stmt
var StmtGetUserByNick *sql.Stmt

func init() {

}

func (f *Forum) scanForum(rows *sql.Rows) error {
	if rows.Next() == true {
		//var slug sql.NullString
		//err := rows.Scan(&f.Posts, slug, &f.Threads, &f.Title, &f.Author)
		//if slug.String != "" {
		//	f.Slug = slug.String
		//}
		err := rows.Scan(&f.Posts, &f.Slug, &f.Threads, &f.Title, &f.Author)
		if err != nil {
			//log.Println("Error in scanForum:", err)
			//log.Println(f)
			return err
		}
		for rows.Next() {
			//var slug sql.NullString
			//err := rows.Scan(&f.Posts, slug, &f.Threads, &f.Title, &f.Author)
			//if slug.String != "" {
			//	f.Slug = slug.String
			//}
			err := rows.Scan(&f.Posts, &f.Slug, &f.Threads, &f.Title, &f.Author)
			if err != nil {
				//log.Println("Error in scanForum:", err)
				//log.Println(f)
				return err
			}
		}
	} else {
		return sql.ErrNoRows
	}
	return nil
}

func CreateForum(db *sql.DB, forum *Forum) error {
	_, err := db.Exec("INSERT INTO forums (slug,title,author) VALUES ($1, $2, $3)", forum.Slug, forum.Title, forum.Author)
	if err != nil {
		//funcname := services.GetFunctionName()
		//log.Printf("Function: %s, Forum %v, Error: %v",funcname ,forum, err)
		return err
	}
	return nil
}

func GetForumBySlug(db *sql.DB, slug string) (*Forum, error) {
	//fmt.Println("GetForum")
	rows,err := StmtGetForumBySlug.Query(slug)
	//rows,err := db.Query("select * from forums where slug = $1", slug)
	defer rows.Close()
	if err != nil {
		//funcname := services.GetFunctionName()
		//log.Printf("Function: %s, Error: %v",funcname , err)
		return nil, err
	}
	forum := &Forum{}
	err = forum.scanForum(rows)
	switch err {
	case sql.ErrNoRows:
		//fmt.Println("GetFourmBySlag ErrNoRows")
		return nil, nil
	case nil:
		return forum, nil
	default:
		//funcname:=services.GetFunctionName()
		//log.Printf("Function: %s, Error: %v",funcname , err)
		return nil, err
	}
	//return forum, nil
}

