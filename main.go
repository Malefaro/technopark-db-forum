package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	_ "github.com/malefaro/technopark-db-forum/routers"
)

//func init(){
//	var filepath string = "resources/init.sql"
//	if _, err := os.Stat(filepath); err == nil {
//		database.Init(filepath)
//	} else {
//		log.Println("file does not exist\n", filepath)
//	}
//}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	defer database.CloseDB()
}
