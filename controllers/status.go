package controllers

import (
	"github.com/astaxie/beego"
	"solo-ci/utils"
)

type StatusController struct {
	beego.Controller
}

func (obj *StatusController) Status() {
	obj.Data["json"] = utils.GetSuccessRender([]string{})
	obj.ServeJSON()
}
