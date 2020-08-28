package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type UserStruct struct {
	Id   int
	Name string `json:"username"`
	Age  int    `json:"age"`
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

type OtherTypeDataContorller struct {
	beego.Controller
}

type FlashController struct {
	beego.Controller
}

// flash
func (f *FlashController) Get() {
	flash := beego.ReadFromRequest(&f.Controller)
	notice := flash.Data["notice"]
	err := flash.Data["error"]
	fmt.Println(err)
	if len(err) != 0 {
		f.TplName = "error.html"
	} else if len(notice) != 0 {
		f.TplName = "success.html"
	} else {
		f.TplName = "flash.html"
	}

}

func (f *FlashController) Post() {

	flash := beego.NewFlash()

	username := f.GetString("username")
	pwd := f.GetString("pwd")
	fmt.Println(username)
	fmt.Println(pwd)

	if len(username) == 0 {
		fmt.Println("username can't null")
		flash.Error("username can't null")
		flash.Store(&f.Controller)
		f.Redirect("/flash_data", 302) //redirect
	} else if pwd != "123456" {
		fmt.Println("pwd worng")
		flash.Error("pwd worng")
		flash.Store(&f.Controller)
		f.Redirect("/flash_data", 302)
	} else {
		flash.Notice("login success")
		flash.Store(&f.Controller)
		f.Redirect("/flash_data", 302)
	}
}

func (o *OtherTypeDataContorller) Get() {
	user := UserStruct{Id: 1, Name: "2131", Age: 18}
	//json
	// o.Data["json"] = &user
	// o.ServeJSON()
	//xml
	// o.Data["xml"] = &user
	// o.ServeXML()
	//jsonp
	// o.Data["jsonp"] = &user
	// o.ServeJSONP()
	// yml
	o.Data["yaml"] = &user
	o.ServeYAML()
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
	// name := p.GetString("username")
	// age, _ := p.GetInt64("age")
	// price, _ := p.GetFloat("price")
	// is_true, _ := p.GetBool("is_true")
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(price)
	// fmt.Println(is_true)

	// id1 := p.Input().Get("username")
	// fmt.Println(id1, "DDDDD")

	// user := UserStruct{}
	// err := p.ParseForm(&user) // notice:binding
	// if err == nil {
	// 	fmt.Println(user)
	// }
	// fmt.Println(err)
	// p.TplName = "success.html"

	// get ajax data
	var user UserStruct
	body := p.Ctx.Input.RequestBody //二进制json数据
	// json.Unmarshal(body, &user)
	fmt.Println(string(body))
	json.Unmarshal(body, &user)
	fmt.Println(user.Age)
	fmt.Println(user)
	result := map[string]string{"code": "200", "message": "处理成功"}
	p.Data["json"] = result
	p.ServeJSON() //返回json格式

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
