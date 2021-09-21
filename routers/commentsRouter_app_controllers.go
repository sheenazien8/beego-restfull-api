package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["example/app/controllers:AuthController"] = append(beego.GlobalControllerRouter["example/app/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example/app/controllers:ExampleController"] = append(beego.GlobalControllerRouter["example/app/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example/app/controllers:StudentController"] = append(beego.GlobalControllerRouter["example/app/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example/app/controllers:StudentController"] = append(beego.GlobalControllerRouter["example/app/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example/app/controllers:StudentController"] = append(beego.GlobalControllerRouter["example/app/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example/app/controllers:StudentController"] = append(beego.GlobalControllerRouter["example/app/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example/app/controllers:StudentController"] = append(beego.GlobalControllerRouter["example/app/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Deletes",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

}
