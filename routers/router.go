// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/malefaro/technopark-db-forum/controllers"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/forum",
			beego.NSInclude(
				&controllers.ForumController{},
			),
		),
		//beego.NSNamespace("/post",
		//	beego.NSInclude(
		//		&controllers.ThreadController2{},
		//	),
		//),
		//beego.NSNamespace("/service",
		//	beego.NSInclude(
		//		&controllers.ThreadController2{},
		//	),
		//),
		//beego.NSNamespace("/thread",
		//	beego.NSInclude(
		//		&controllers.ThreadController2{},
		//	),
		//),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.Router("/", &controllers.UserController{}, "get:GetAll")
	//beego.Router("/q", &controllers.Test{}, "get:Test")
}