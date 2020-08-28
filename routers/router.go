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
	beego.Router("/params", &controllers.ParamsController{})
}
