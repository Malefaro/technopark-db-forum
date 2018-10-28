package controllers

import (
	"encoding/json"
	"github.com/malefaro/technopark-db-forum/database"
	"github.com/malefaro/technopark-db-forum/models"

	"net/http"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Post
// @Description create user
// @Param	nickname		path 	string	true	"nickname from uri"
// @Param profile body models.User true "profile"
// @Success 201 {object} models.User
// @Failure 403 :uid is empty
// @router /:nickname/create [post]
func (u *UserController) Post() {
	db := database.GetDataBase()
	body := u.Ctx.Input.RequestBody
	nickname := u.GetString(":nickname")
	user := &models.User{Nickname: nickname}
	json.Unmarshal(body, &user)
	result := make([]*models.User, 1)
	var us *models.User
	if us, _ = models.GetUserByEmail(db, user.Email); us != nil {
		result = append(result, us)
	}
	if us, _ = models.GetUserByNickname(db, user.Nickname); us != nil {
		if len(result) == 0 || result[0] != us {
			result = append(result, us)
		}
	}
	if len(result) != 0 {
		u.Ctx.Output.SetStatus(http.StatusConflict)
		u.Data["json"] = result
		u.ServeJSON()
		return
	}
	//fmt.Println(u.DB)
	//fmt.Println(user)
	models.CreateUser(db, user)
	u.Data["json"] = user
	u.Ctx.Output.SetStatus(http.StatusCreated)
	u.ServeJSON()
}
