package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"github.com/satori/go.uuid"
	"solo-ci/interfaces"
	"time"
	"go/build"
)

type Project struct {
	Id          int `orm:"pk;auto;unique" json:"id"`
	ProjectId   string `json:"project_id"`           //uuid
	Name        string `json:"name" form:"name"`     //name
	Type        string `json:"type" form:"type"`     //github,gitlab
	Url         string `json:"url" form:"url"`       //仓库地址
	Path        string `json:"path" form:"path"`     //file 地址
	Branch      string `json:"branch" form:"branch"` //分支
	SecretToken string `json:"secret_token" form:"secret_token"`
	Build       []*Build `orm:"reverse(many)" json:"-"`
}

// return  id,err
func (obj *Project) Add() (string, error) {
	u := uuid.NewV1()
	obj.ProjectId = u.String()
	o := orm.NewOrm()
	qs := o.QueryTable("project")
	qs = qs.Filter("name", obj.Name)
	if count, err := qs.Count(); count != 0 {
		if err == nil {
			err = errors.New("Already Exist")
		}
		return "", err
	} else {
		_, err := o.Insert(obj)
		if err != nil {
			return "", err
		} else {
			return obj.ProjectId, nil
		}
	}
}

//return id,err
func (obj *Project) Delete() (error) {
	o := orm.NewOrm()
	_, err := o.Delete(obj)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//return isSuccess,err
func (obj *Project) Update() (error) {
	o := orm.NewOrm()
	_, err := o.Update(obj)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//return obj list,err
func (obj *Project) Get() (error) {
	o := orm.NewOrm()
	qs := o.QueryTable(obj)
	err := qs.One(obj)
	if err != nil {
		return err
	} else {
		return nil
	}
	return nil
}