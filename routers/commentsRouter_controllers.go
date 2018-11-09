package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/:slug/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"],
        beego.ControllerComments{
            Method: "Details",
            Router: `/:slug/details`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"],
        beego.ControllerComments{
            Method: "Threads",
            Router: `/:slug/threads`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:ForumController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
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

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:UserController"],
        beego.ControllerComments{
            Method: "ProfileGet",
            Router: `/:nickname/profile`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/malefaro/technopark-db-forum/controllers:UserController"],
        beego.ControllerComments{
            Method: "ProfilePost",
            Router: `/:nickname/profile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}