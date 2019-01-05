package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/models"
	_ "github.com/malefaro/technopark-db-forum/routers"
	"log"
	"os"
	"runtime"
)

func init(){
	var filepath string = "resources/init.sql"
	if _, err := os.Stat(filepath); err == nil {
		database.Init(filepath)
	} else {
		log.Println("file does not exist\n", filepath)
	}
	db := database.GetDataBase()
	var err error
	//fmt.Println("INIT stmt")
	models.StmtGetForumBySlug, err = db.Prepare("select * from forums where slug = $1")
	if err != nil {
		fmt.Println("error while preparing", err)
		return
	}
	models.StmtGetThreadByID, err = db.Prepare("select * from threads where id = $1")
	if err != nil {
		fmt.Println("error while preparing", err)
		return
	}
	models.StmtGetThreadBySlug, err = db.Prepare("select * from threads where slug = $1")
	if err != nil {
		fmt.Println("error while preparing", err)
		return
	}
	models.StmtGetUserByNick, err = db.Prepare("select * from users where nickname = $1")
	if err != nil {
		fmt.Println("error while preparing", err)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.Log.AccessLogs = false
	//beego.BeeLogger.Close()
	//beego.BeeLogger.Write([]byte("BLYA"))
	beego.Run()

	defer database.CloseDB()
}
