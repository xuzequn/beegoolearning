package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type RouteController struct {
	beego.Controller
}

func (r *RouteController) Get() {
	if r.Ctx.Request.Method == "POST" {
		fmt.Printf("111111")
		// r.StopRun()
		r.TplName = "test_router_post.html"
	} else if r.Ctx.Request.Method == "GET" {
		// id := r.Ctx.Input.Param(":id")
		// fmt.Println("===============")
		// fmt.Println(id)
		r.TplName = "test_router.html"
	}
}
