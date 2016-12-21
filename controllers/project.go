package controllers

import (
	"github.com/astaxie/beego"
	"solo-ci/models"
	"solo-ci/utils"
	"fmt"
	"errors"
	"github.com/astaxie/beego/orm"
	"debug/elf"
	"io/ioutil"
	"encoding/json"
)

type ProjectController struct {
	beego.Controller
}

func (obj *ProjectController) Add() {
	project := new(models.Project)
	if err := obj.ParseForm(project); err != nil {
		obj.Data["json"] = utils.GetClientErrRender()
	} else {
		id, err := project.Add()
		if err != nil {
			obj.Data["json"] = utils.GetErrorRender(err.Error(), 400)
		} else {
			obj.Data["json"] = utils.GetSuccessRender(map[string]string{"project_id":id})
		}
	}
	obj.ServeJSON()
}

func (obj *ProjectController) Delete() {
	projectId := obj.Ctx.Input.Param(":project_id")
	fmt.Println(projectId)
	project := &models.Project{ProjectId:projectId}
	err := project.Delete()
	if err != nil {
		obj.Data["json"] = utils.GetErrorRender(err.Error(), 400)
	} else {
		obj.Data["json"] = utils.GetSuccessRender(nil)
	}
	obj.ServeJSON()
}

func (obj *ProjectController) Update() {
	projectId := obj.Ctx.Input.Param(":project_id")
	project := &models.Project{ProjectId:projectId}
	if err := obj.ParseForm(project); err != nil {
		obj.Data["json"] = utils.GetClientErrRender()
	} else {
		err := project.Update()
		if err != nil {
			obj.Data["json"] = utils.GetErrorRender(err.Error(), 400)
		} else {
			obj.Data["json"] = utils.GetSuccessRender(nil)
		}
	}
	obj.ServeJSON()
}

func (obj *ProjectController) Get() {
	projectId := obj.Ctx.Input.Param(":project_id")
	project := &models.Project{ProjectId:projectId}
	err := project.Get()
	if err != nil {
		obj.Data["json"] = utils.GetErrorRender(err.Error(), 400)
	} else {
		obj.Data["json"] = utils.GetSuccessRender(project)
	}
	obj.ServeJSON()
}

func (obj *ProjectController) WebHook() {
	//告诉git 接受成功
	obj.Data["json"] = utils.GetSuccessRender(nil)
	obj.ServeJSON()
	//执行脚本
	project := new(models.Project)
	project.ProjectId = obj.Ctx.Input.Param(":project_id")
	o := orm.NewOrm()
	if err := o.Read(project); err != nil && project.Name == "" {
		beego.Info("This object", project.ProjectId, "not exist")
		return
	}
	switch project.Type {
	case "gitlab":
		gitlabHook := new(models.GitlabHook)
		bodyMsg, _ := ioutil.ReadAll(obj.Ctx.Request.Body)
		json.Unmarshal(bodyMsg, gitlabHook)
		if project.SecretToken != "" && obj.Ctx.Request.Header.Get("X-Gitlab-Token") != project.SecretToken {
			beego.Info(project.ProjectId,"Secret token error")
			return
		}
		build := models.NewBuild(project)
		o := orm.NewOrm()
		o.Insert(build)
	case "github":
		beego.Info("This type will support in next version")
	case "bitbucket":
		beego.Info("This type will support in next version")
	default:
		beego.Info("Don't have this type")
	}
}