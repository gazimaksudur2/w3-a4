package main

import (
	_ "w3-a4/routers"

	beego "github.com/beego/beego/v2/server/web"

	// _ "w3-a4/docs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
