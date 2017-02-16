package main

import (
	"os"
	"os/exec"
	"github.com/astaxie/beego"
	"solo-ci/conf"
	_ "solo-ci/routers"
)

func main() {
	//创建工作区
	if _, err := os.Stat("workspace"); os.IsNotExist(err) {
		os.Mkdir("workspace", 0766)
	}
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0766)
	}
	//检查git
	_, err := exec.Command("git", "--version").CombinedOutput()
	if err != nil {
		beego.Info("Git not Install")
	}
	//检查golang
	conf.GOPATH = os.Getenv("GOPATH")
	conf.GOROOT = os.Getenv("GOROOT")
	if conf.GOPATH == "" {
		beego.Info("GOPATH not set")
	}
	if conf.GOROOT == "" {
		beego.Info("GOROOT not set")
	}
	beego.Info("check success!")
	beego.Run()
}