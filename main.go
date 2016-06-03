package main

import (
	"flag"

	_ "parkinglots/docs"
	_ "parkinglots/etc"
	_ "parkinglots/routers"

	"github.com/astaxie/beego"
)


var (
	user = flag.String("user", "root", "MySQL username")
	pwd  = flag.String("pwd", "", "MySQL password")
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
