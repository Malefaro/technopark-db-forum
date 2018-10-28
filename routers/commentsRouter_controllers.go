package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ThreadController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/:slug_or_id/create`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ThreadController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/:nickname/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
