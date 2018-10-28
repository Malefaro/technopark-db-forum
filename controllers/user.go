package controllers

import (
	"database/sql"
	"fmt"
	"github.com/malefaro/technopark-db-forum/models"
	"log"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
	DB *sql.DB
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:nickname/create [post]
func (u *UserController) Post() {
	nickname := u.GetString(":nickname")
	name := u.Ctx.Input.Query("name")
	log.Println("GET NAME FROM QUERY",name)

	//if nickname != "" {
	//	user, err := models.GetUser(uid)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = user
	//	}
	//}
	user := &models.User{"test user","testemail@gmail.com","keker",nickname}

	fmt.Println(u.DB)
	fmt.Println(user)
	models.CreateUser(u.DB,user)
	u.Data["json"] = user
	u.ServeJSON()
}


