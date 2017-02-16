package controllers

import (
	"github.com/astaxie/beego"
	"solo-ci/models"
	"solo-ci/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"encoding/json"
	"net/http/httputil"
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

func (obj *ProjectController) GetList() {
	page, _ := obj.GetInt("page", 0)
	pageSize, _ := obj.GetInt("pageSize", 20)
	obj.Data["json"] = utils.GetSuccessRender(models.GetList(page, pageSize))
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
	if err := o.Read(project, "project_id"); err != nil && project.Name == "" {
		beego.Info("This object", project.ProjectId, "not exist")
		return
	}
	switch project.Type {
	case "gitlab":
		gitlabHook := new(models.GitlabHook)
		bodyMsg, _ := ioutil.ReadAll(obj.Ctx.Request.Body)
		json.Unmarshal(bodyMsg, gitlabHook)
		if gitlabHook.Ref != "refs/heads/" + project.Branch {
			beego.Info("Branch not same")
			return
		}
		if project.SecretToken != "" && obj.Ctx.Request.Header.Get("X-Gitlab-Token") != project.SecretToken {
			beego.Info(project.ProjectId, "Secret token error")
			return
		}
		go models.NewBuild(project)
	case "github":
		githubHook := new(models.GithubHook)
		bodyMsg, _ := ioutil.ReadAll(obj.Ctx.Request.Body)
		jjj, _ := httputil.DumpRequest(obj.Ctx.Request, true)
		json.Unmarshal(bodyMsg, githubHook)
		beego.Info(string(jjj))
		if githubHook.Ref != "refs/heads/" + project.Branch {
			beego.Info("Branch not same")
			return
		}
		if project.SecretToken != "" && obj.Ctx.Request.Header.Get("X-Hub-Signature") != project.SecretToken {
			beego.Info(project.ProjectId, "Secret token error")
			return
		}
		go models.NewBuild(project)
	default:
		beego.Info("Don't have this type")
	}
}