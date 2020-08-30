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
	beego.Router("/params/?:id", &controllers.ParamsController{}, "get:GET")
	beego.Router("/params", &controllers.ParamsController{})
	beego.Router("/other_type_data", &controllers.OtherTypeDataContorller{})
	beego.Router("/flash_data", &controllers.FlashController{})
}
