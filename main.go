package main

import (
	_ "beegolearning/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {

	httpport := beego.AppConfig.String("httpport")
	fmt.Printf(httpport)
	beego.Run()
}
