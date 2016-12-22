package main

import (
	"github.com/astaxie/beego"
	"os"
	_ "solo-ci/routers"
	"solo-ci/conf"
	"os/exec"
)

func main() {
	//创建工作区
	if _, err := os.Stat("workspace"); os.IsNotExist(err) {
		os.Mkdir("workspace", 0766)
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