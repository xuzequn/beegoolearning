package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserStruct struct {
	Id   int
	Name string `form:"username"`
	Age  int    `form:"age"`
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

//前端传参数到后端

type ParamsController struct {
	beego.Controller
}

func (p *ParamsController) Get() {
	name := p.GetString("name")
	// url: /params
	fmt.Println("===========================")
	fmt.Println(name)
	fmt.Println("===========================")
	name2 := p.Input().Get("name")
	fmt.Println(name2)
	fmt.Println("===========================")
	// name3 := p.Ctx.Input.Param("name")
	// fmt.Println(name3)

	// url: /params/:id:int
	id := p.GetString(":id")
	fmt.Println(id)

	// id1 := p.Input().Get(":id")
	// fmt.Println(id1, "DDDDD")

	id2 := p.Ctx.Input.Param(":id")
	fmt.Println(id2)

	p.TplName = "param_test.html"

}

func (p *ParamsController) Post() {
	name := p.GetString("username")
	age, _ := p.GetInt64("age")
	// price, _ := p.GetFloat("price")
	// is_true, _ := p.GetBool("is_true")
	fmt.Println(name)
	fmt.Println(age)
	// fmt.Println(price)
	// fmt.Println(is_true)

	// id1 := p.Input().Get("username")
	// fmt.Println(id1, "DDDDD")

	user := UserStruct{}
	err := p.ParseForm(&user) // notice:binding
	if err == nil {
		fmt.Println(user)
	}
	fmt.Println(err)
	p.TplName = "success.html"
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
