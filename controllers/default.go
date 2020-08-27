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

type StaticController struct {
	beego.Controller
}

func (s *StaticController) Get() {

	s.TplName = "static_test.html"

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

	// map
	test_map := map[string]string{"name": "sssss", "age": "19", "addr": "ssafdsf"}
	u.Data["test_map"] = test_map

	// map_struct
	test_map_struct := make(map[int]UserStruct)
	test_map_struct[1] = UserStruct{Id: 1, Name: "sssss", Age: 19}
	test_map_struct[2] = UserStruct{Id: 2, Name: "sssss", Age: 19}
	test_map_struct[3] = UserStruct{Id: 3, Name: "sssss", Age: 19}
	test_map_struct[4] = UserStruct{Id: 4, Name: "sssss", Age: 19}
	u.Data["map_struct"] = test_map_struct

	// sclice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	u.Data["slice"] = slice

	u.TplName = "user.html"

}
