package main

import (
	"github.com/astaxie/beego"
	_ "solo-ci/models"
	_ "solo-ci/routers"
)

func main() {
	beego.Run()
}