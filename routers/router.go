package routers

import (
	"beegolearning/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/test_static", &controllers.StaticController{})
	// beego.Router("/params/?:id:int", &controllers.ParamsController{})

	// beego.Router("/params/?:id", &controllers.ParamsController{}, "get:GET")
	// 自动路由
	// beego.AutoRouter(&controllers.ParamsController{})
	// 自定义路由
	// beego.Router("/params/:id([0-9]+)", &controllers.ParamsController{}, "get:Get1")
	beego.Router("/params/:id([0-9]+)", &controllers.ParamsController{}, "get,post:Get")
	beego.Router("/other_type_data", &controllers.OtherTypeDataContorller{})
	beego.Router("/flash_data", &controllers.FlashController{})
	//几种路由参数
	//1固定路由
	// beego.Router("/test_router", &controllers.RouteController{})
	//2正则路由
	// beego.Router("/test_router/?:id:int", &controllers.RouteController{})
	//3自动路由
	// beego.AutoRouter(&controllers.RouteController{})
	//4自定义路由
	beego.Router("/test_router/?:id([0-9]+)", &controllers.RouteController{}, "get,post:Get")
}
