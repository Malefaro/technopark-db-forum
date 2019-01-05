package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/malefaro/technopark-db-forum/database"
	_ "github.com/malefaro/technopark-db-forum/routers"
	"runtime"
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
