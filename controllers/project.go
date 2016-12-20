package controllers

import (
	"github.com/astaxie/beego"
	"solo-ci/models"
	"solo-ci/utils"
	"fmt"
)

type ProjectController struct {
	beego.Controller
}

func (obj *ProjectController) Add() {
	project := new(models.Project)
	if err := obj.ParseForm(project); err != nil {
		obj.Data["json"] = utils.GetClientErrRender()
		obj.ServeJSON()
	} else {
		id, err := project.Add()
		if err != nil {
			obj.Data["json"] = utils.GetErrorRender(err.Error(), 400)
			obj.ServeJSON()
		} else {
			obj.Data["json"] = utils.GetSuccessRender(map[string]string{"project_id":id})
			obj.ServeJSON()
		}
	}
}

func (obj *ProjectController) Delete() {
	projectId := obj.Ctx.Input.Param(":project_id")
	fmt.Println(projectId)
	project := &models.Project{ProjectId:projectId}
	err := project.Delete()
	if err != nil {
		obj.Data["json"] = utils.GetErrorRender(err.Error(),400)
		obj.ServeJSON()
	} else {
		obj.Data["json"] = utils.GetSuccessRender(nil)
		obj.ServeJSON()
	}
}

func (obj *ProjectController) Update() {
	projectId := obj.Ctx.Input.Param(":project_id")
	project := &models.Project{ProjectId:projectId}
	if err := obj.ParseForm(project); err != nil {
		obj.Data["json"] = utils.GetClientErrRender()
		obj.ServeJSON()
	} else {
		err := project.Update()
		if err != nil {
			obj.Data["json"] = utils.GetErrorRender(err.Error(),400)
			obj.ServeJSON()
		} else {
			obj.Data["json"] = utils.GetSuccessRender(nil)
			obj.ServeJSON()
		}
	}
	project.Update()
}

func (obj *ProjectController) Get() {
	projectId := obj.Ctx.Input.Param(":project_id")
	project := &models.Project{ProjectId:projectId}
	err := project.Get()
	if err != nil {
		obj.Data["json"] = utils.GetErrorRender(err.Error(),400)
	} else {
		obj.Data["json"] = utils.GetSuccessRender(project)
	}
	obj.ServeJSON()
}

func (obj *ProjectController) WebHook() {

}