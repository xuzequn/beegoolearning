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
	u.TplName = "user.html"

	// array

}
