package controllers

import (
	"github.com/astaxie/beego"
)

type UserStruct struct {
	Id   int
	Name string
	Age  int
}

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (u *UserController) Get() {

	// struct
	user := UserStruct{Id: 1, Name: "ssss", Age: 18}
	u.Data["user"] = user

	// array
	arr := [5]int{9, 8, 7, 6, 5}
	u.Data["arr"] = arr

	//arr_struct
	arr_struct := [3]UserStruct{{Id: 1, Name: "2222"}, {Id: 1, Name: "2222"}, {Id: 1, Name: "2222"}}
	u.Data["arr_struct"] = arr_struct

	u.TplName = "user.html"

}
